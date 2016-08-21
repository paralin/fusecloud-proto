package rql

import (
	"github.com/fuserobotics/proto/region"
	r "gopkg.in/dancannon/gorethink.v2"
)

var RegionsTable string = "regions"

func FetchAllRegions(rctx *r.Session) ([]*region.Region, error) {
	var result []*region.Region

	res, err := r.DB("global").Table(RegionsTable).Run(rctx)
	if err == nil {
		err = res.All(&result)
	}
	if err != nil {
		res.Close()
		return nil, err
	}

	return result, res.Close()
}

func FetchRegionById(id string, rctx *r.Session) (*region.Region, error) {
	result := &region.Region{}

	res, err := r.DB("global").Table(RegionsTable).Get(id).Run(rctx)
	if err == nil {
		err = res.One(&result)
	}
	if err != nil {
		res.Close()
		return nil, err
	}

	return result, res.Close()
}

func CreateRegion(rctx *r.Session, reg *region.Region) error {
	_, err := r.DB("global").
		Table(RegionsTable).
		Insert(reg).
		RunWrite(rctx)
	return err
}

func UpdateRegion(rctx *r.Session, region string, update map[string]interface{}) error {
	_, err := r.DB("global").
		Table(RegionsTable).
		Get(region).
		Update(update).
		RunWrite(rctx)
	return err
}
