package rql

import (
	"github.com/fuserobotics/proto/device"
	r "gopkg.in/dancannon/gorethink.v2"
)

func GetDatabaseByRegion(region string) string {
	return "region_" + region
}

func GetDeviceByHostname(rctx *r.Session, hostname, region string) (*device.Device, error) {
	device := &device.Device{}
	cursor, err := r.DB(GetDatabaseByRegion(region)).Table("devices").Get(hostname).Run(rctx)
	defer cursor.Close()
	if err != nil {
		return nil, err
	}
	return device, cursor.One(device)
}
