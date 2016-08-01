package device_cql

import (
	"errors"
	"fmt"

	"github.com/paralin/cqlpb"
	"github.com/paralin/cqlpb/marshal"
	"github.com/relops/cqlr"
	"github.com/synrobo/cassandra"
	"github.com/synrobo/proto/device"
)

func GetDevices(ctx *cassandra.CassandraContext) ([]*device.Device, error) {
	var result []*device.Device

	b := cqlpb.BindQuery(ctx.Session.Query(fmt.Sprintf("SELECT * FROM %s", device.DevicesTable)))

	var itr *device.Device
	for {
		itr = &device.Device{}
		hasMore, err := b.Scan(itr)

		if err != nil {
			defer b.Close()
			return nil, err
		}

		if !hasMore {
			break
		}

		result = append(result, itr)
	}

	if result == nil {
		result = []*device.Device{}
	}

	return result, b.Close()
}

func GetDeviceByHostname(ctx *cassandra.CassandraContext, id string) (error, *device.Device) {
	b := cqlpb.BindQuery(ctx.Session.Query(fmt.Sprintf("SELECT * FROM %s WHERE hostname = ? LIMIT 1", device.DevicesTable), id))
	defer b.Close()

	nreg := &device.Device{}
	succ, err := b.Scan(nreg)
	if err != nil {
		return err, nil
	}
	if !succ {
		return errors.New("Not found."), nil
	}
	return nil, nreg
}

// Probably able to make this dynamically?
type insertDeviceStruct struct {
	Hostname string
	Region   string
	Proto    []byte
	Fqdn     string
}

func CreateDevice(ctx *cassandra.CassandraContext, ndev *device.Device) error {
	toInsert, err := marshal.Marshal(ndev, device.DevicesTableTemplate)
	if err != nil {
		return err
	}

	tins := &insertDeviceStruct{
		Hostname: toInsert["hostname"].(string),
		Region:   toInsert["region"].(string),
		Proto:    toInsert["proto"].([]byte),
		Fqdn:     toInsert["fqdn"].(string),
	}

	q := cqlr.Bind(fmt.Sprintf("INSERT INTO %s (hostname, region, proto, fqdn) VALUES (?, ?, ?, ?)", device.DevicesTable), tins)
	return q.Exec(ctx.Session)
}
