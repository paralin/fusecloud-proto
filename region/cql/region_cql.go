package region_cql

import (
	"errors"
	"fmt"

	"github.com/paralin/cqlpb"
	"github.com/paralin/cqlpb/marshal"
	"github.com/relops/cqlr"
	"github.com/synrobo/cassandra"
	"github.com/synrobo/proto/region"
)

func GetAllRegions(ctx *cassandra.CassandraContext) (error, []*region.Region) {
	var result []*region.Region

	b := cqlpb.BindQuery(ctx.Session.Query(fmt.Sprintf("SELECT * FROM %s", region.RegionTable)))

	var itr *region.Region
	for {
		itr = &region.Region{}
		hasMore, err := b.Scan(itr)

		if err != nil {
			return err, nil
		}

		if !hasMore {
			break
		}

		result = append(result, itr)
	}

	if result == nil {
		result = []*region.Region{}
	}

	return b.Close(), result
}

func GetRegionById(ctx *cassandra.CassandraContext, id string) (error, *region.Region) {
	b := cqlpb.BindQuery(ctx.Session.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = ? LIMIT 1", region.RegionTable), id))

	nreg := &region.Region{}
	succ, err := b.Scan(nreg)
	if err == nil {
		err = b.Close()
	} else {
		defer b.Close()
	}
	if err != nil {
		return err, nil
	}
	if !succ {
		return errors.New("Not found."), nil
	}
	return nil, nreg
}

// Probably able to make this dynamically?
type insertRegionStruct struct {
	Id    string
	Name  string
	Proto []byte
}

func CreateRegion(ctx *cassandra.CassandraContext, nreg *region.Region) error {
	toInsert, err := marshal.Marshal(nreg, region.RegionTableTemplate)
	if err != nil {
		return err
	}

	tins := &insertRegionStruct{
		Id:    toInsert["id"].(string),
		Name:  toInsert["name"].(string),
		Proto: toInsert["proto"].([]byte),
	}

	q := cqlr.Bind(fmt.Sprintf("INSERT INTO %s (id, name, proto) VALUES (?, ?, ?)", region.RegionTable), tins)
	return q.Exec(ctx.Session)
}

func BuildContextForRegion(cctx *cassandra.CassandraContext, cassandraIp string, regionId string) (*cassandra.CassandraContext, *region.Region, error) {
	err, reg := GetRegionById(cctx, regionId)
	if err != nil {
		return nil, nil, err
	}

	ctx, err := BuildCassandraContextForMetaRegion(reg, cassandraIp)
	if err != nil {
		return nil, reg, err
	}
	return ctx, reg, nil
}
