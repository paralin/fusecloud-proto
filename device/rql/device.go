package rql

import (
	"github.com/fuserobotics/proto/device"
	region_rql "github.com/fuserobotics/proto/region/rql"
	r "gopkg.in/dancannon/gorethink.v2"
)

var DevicesTable string = "devices"

func GetDeviceByHostname(rctx *r.Session, hostname, region string) (*device.Device, error) {
	device := &device.Device{}
	cursor, err := r.DB(region_rql.GetDatabaseByRegion(region)).Table(DevicesTable).Get(hostname).Run(rctx)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	if cursor.IsNil() {
		return nil, nil
	}
	return device, cursor.One(device)
}

func GetAllDevices(rctx *r.Session, region string) ([]*device.Device, error) {
	var res []*device.Device
	cursor, err := r.DB(region_rql.GetDatabaseByRegion(region)).Table(DevicesTable).Run(rctx)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	return res, cursor.All(&res)
}

func CreateDevice(rctx *r.Session, dev *device.Device) error {
	_, err := r.DB(region_rql.GetDatabaseByRegion(dev.Region)).
		Table(DevicesTable).
		Insert(dev).
		RunWrite(rctx)
	return err
}

func UpdateDevice(rctx *r.Session, hostname, region string, update map[string]interface{}) error {
	_, err := r.DB(region_rql.GetDatabaseByRegion(region)).
		Table(DevicesTable).
		Get(hostname).
		Update(update).
		RunWrite(rctx)
	return err
}
