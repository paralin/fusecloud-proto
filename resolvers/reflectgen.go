// +build codegen

package main

import (
	"github.com/fuserobotics/proto/auth"
	"github.com/fuserobotics/proto/common"
	"github.com/fuserobotics/proto/device"
	"github.com/fuserobotics/proto/region"
	"github.com/fuserobotics/proto/timesync"
)

type GraphQLProto interface {
	GraphQLTypeName() string
}

// Map of GraphQL simplified names to proto names.
var AllGraphQLTypes = map[string]GraphQLProto{
	"auth.UserPrivateData": &auth.UserPrivateData{},
	"auth.UserPrivateData_ScryptParams": &auth.UserPrivateData_ScryptParams{},
	"auth.UserPublicIdentity": &auth.UserPublicIdentity{},
	"common.CertChain": &common.CertChain{},
	"common.GeoLocation": &common.GeoLocation{},
	"common.IPAddress": &common.IPAddress{},
	"common.IPRange": &common.IPRange{},
	"device.Device": &device.Device{},
	"device.DeviceNetworkTemplate": &device.DeviceNetworkTemplate{},
	"device.DeviceWithStatus": &device.DeviceWithStatus{},
	"device.Device_DeviceIdentity": &device.Device_DeviceIdentity{},
	"device.Device_DeviceInterfaceConfig": &device.Device_DeviceInterfaceConfig{},
	"device.Device_DeviceInterfaceConfig_WifiConfig": &device.Device_DeviceInterfaceConfig_WifiConfig{},
	"device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork": &device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork{},
	"device.Device_DeviceNetworkSettings": &device.Device_DeviceNetworkSettings{},
	"device.KVGDeviceSubKeys": &device.KVGDeviceSubKeys{},
	"region.Region": &region.Region{},
	"timesync.GetUnixTimeResponse": &timesync.GetUnixTimeResponse{},
}