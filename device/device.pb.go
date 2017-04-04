// Code generated by protoc-gen-go.
// source: github.com/fuserobotics/proto/device/device.proto
// DO NOT EDIT!

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	github.com/fuserobotics/proto/device/device.proto

It has these top-level messages:
	Device
	DeviceNetworkTemplate
*/
package device

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/fuserobotics/proto/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceStatus int32

const (
	DeviceStatus_UNKNOWN DeviceStatus = 0
	DeviceStatus_ALIVE   DeviceStatus = 1
)

var DeviceStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "ALIVE",
}
var DeviceStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"ALIVE":   1,
}

func (x DeviceStatus) String() string {
	return proto.EnumName(DeviceStatus_name, int32(x))
}
func (DeviceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType int32

const (
	// Standard wifi access point
	Device_DeviceInterfaceConfig_WifiConfig_AP    Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType = 0
	Device_DeviceInterfaceConfig_WifiConfig_ADHOC Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType = 1
)

var Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_name = map[int32]string{
	0: "AP",
	1: "ADHOC",
}
var Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_value = map[string]int32{
	"AP":    0,
	"ADHOC": 1,
}

func (x Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType) String() string {
	return proto.EnumName(Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_name, int32(x))
}
func (Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0, 0}
}

// Device: a full-linux computer running serf, docker stack
type Device struct {
	Hostname        string                        `protobuf:"bytes,1,opt,name=hostname" json:"hostname"`
	Region          string                        `protobuf:"bytes,2,opt,name=region" json:"region"`
	NetworkSettings *Device_DeviceNetworkSettings `protobuf:"bytes,3,opt,name=network_settings,json=networkSettings" json:"network_settings"`
	Identity        *Device_DeviceIdentity        `protobuf:"bytes,4,opt,name=identity" json:"identity"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Device) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Device) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Device) GetNetworkSettings() *Device_DeviceNetworkSettings {
	if m != nil {
		return m.NetworkSettings
	}
	return nil
}

func (m *Device) GetIdentity() *Device_DeviceIdentity {
	if m != nil {
		return m.Identity
	}
	return nil
}

type Device_DeviceNetworkSettings struct {
	Ip        *common.IPAddress               `protobuf:"bytes,1,opt,name=ip" json:"ip"`
	Interface []*Device_DeviceInterfaceConfig `protobuf:"bytes,2,rep,name=interface" json:"interface"`
}

func (m *Device_DeviceNetworkSettings) Reset()                    { *m = Device_DeviceNetworkSettings{} }
func (m *Device_DeviceNetworkSettings) String() string            { return proto.CompactTextString(m) }
func (*Device_DeviceNetworkSettings) ProtoMessage()               {}
func (*Device_DeviceNetworkSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Device_DeviceNetworkSettings) GetIp() *common.IPAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Device_DeviceNetworkSettings) GetInterface() []*Device_DeviceInterfaceConfig {
	if m != nil {
		return m.Interface
	}
	return nil
}

type Device_DeviceInterfaceConfig struct {
	Devname      string                                   `protobuf:"bytes,1,opt,name=devname" json:"devname"`
	Ip           *common.IPAddress                        `protobuf:"bytes,2,opt,name=ip" json:"ip"`
	GatewayIp    *common.IPAddress                        `protobuf:"bytes,3,opt,name=gateway_ip,json=gatewayIp" json:"gateway_ip"`
	WifiSettings *Device_DeviceInterfaceConfig_WifiConfig `protobuf:"bytes,4,opt,name=wifi_settings,json=wifiSettings" json:"wifi_settings"`
	Dns          []*common.IPAddress                      `protobuf:"bytes,5,rep,name=dns" json:"dns"`
}

func (m *Device_DeviceInterfaceConfig) Reset()                    { *m = Device_DeviceInterfaceConfig{} }
func (m *Device_DeviceInterfaceConfig) String() string            { return proto.CompactTextString(m) }
func (*Device_DeviceInterfaceConfig) ProtoMessage()               {}
func (*Device_DeviceInterfaceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Device_DeviceInterfaceConfig) GetDevname() string {
	if m != nil {
		return m.Devname
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig) GetIp() *common.IPAddress {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Device_DeviceInterfaceConfig) GetGatewayIp() *common.IPAddress {
	if m != nil {
		return m.GatewayIp
	}
	return nil
}

func (m *Device_DeviceInterfaceConfig) GetWifiSettings() *Device_DeviceInterfaceConfig_WifiConfig {
	if m != nil {
		return m.WifiSettings
	}
	return nil
}

func (m *Device_DeviceInterfaceConfig) GetDns() []*common.IPAddress {
	if m != nil {
		return m.Dns
	}
	return nil
}

type Device_DeviceInterfaceConfig_WifiConfig struct {
	Network      []*Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork `protobuf:"bytes,1,rep,name=network" json:"network"`
	ExtraOptions string                                                 `protobuf:"bytes,2,opt,name=extra_options,json=extraOptions" json:"extra_options"`
}

func (m *Device_DeviceInterfaceConfig_WifiConfig) Reset() {
	*m = Device_DeviceInterfaceConfig_WifiConfig{}
}
func (m *Device_DeviceInterfaceConfig_WifiConfig) String() string { return proto.CompactTextString(m) }
func (*Device_DeviceInterfaceConfig_WifiConfig) ProtoMessage()    {}
func (*Device_DeviceInterfaceConfig_WifiConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0}
}

func (m *Device_DeviceInterfaceConfig_WifiConfig) GetNetwork() []*Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Device_DeviceInterfaceConfig_WifiConfig) GetExtraOptions() string {
	if m != nil {
		return m.ExtraOptions
	}
	return ""
}

type Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork struct {
	Ssid         string                                                  `protobuf:"bytes,1,opt,name=ssid" json:"ssid"`
	NetType      Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType `protobuf:"varint,9,opt,name=net_type,json=netType,enum=device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType" json:"net_type"`
	Psk          string                                                  `protobuf:"bytes,6,opt,name=psk" json:"psk"`
	PskEncoded   bool                                                    `protobuf:"varint,10,opt,name=psk_encoded,json=pskEncoded" json:"psk_encoded"`
	Proto        string                                                  `protobuf:"bytes,2,opt,name=proto" json:"proto"`
	KeyMgmt      string                                                  `protobuf:"bytes,3,opt,name=key_mgmt,json=keyMgmt" json:"key_mgmt"`
	Pairwise     string                                                  `protobuf:"bytes,4,opt,name=pairwise" json:"pairwise"`
	Group        string                                                  `protobuf:"bytes,5,opt,name=group" json:"group"`
	Frequency    string                                                  `protobuf:"bytes,8,opt,name=frequency" json:"frequency"`
	ExtraOptions string                                                  `protobuf:"bytes,7,opt,name=extra_options,json=extraOptions" json:"extra_options"`
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) Reset() {
	*m = Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork{}
}
func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) String() string {
	return proto.CompactTextString(m)
}
func (*Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) ProtoMessage() {}
func (*Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0, 0}
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetNetType() Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType {
	if m != nil {
		return m.NetType
	}
	return Device_DeviceInterfaceConfig_WifiConfig_AP
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetPsk() string {
	if m != nil {
		return m.Psk
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetPskEncoded() bool {
	if m != nil {
		return m.PskEncoded
	}
	return false
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetKeyMgmt() string {
	if m != nil {
		return m.KeyMgmt
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetPairwise() string {
	if m != nil {
		return m.Pairwise
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetFrequency() string {
	if m != nil {
		return m.Frequency
	}
	return ""
}

func (m *Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork) GetExtraOptions() string {
	if m != nil {
		return m.ExtraOptions
	}
	return ""
}

type Device_DeviceIdentity struct {
	// Generated certs, latest is first
	Chain     []*common.CertChain `protobuf:"bytes,2,rep,name=chain" json:"chain"`
	PublicKey string              `protobuf:"bytes,3,opt,name=public_key,json=publicKey" json:"public_key"`
}

func (m *Device_DeviceIdentity) Reset()                    { *m = Device_DeviceIdentity{} }
func (m *Device_DeviceIdentity) String() string            { return proto.CompactTextString(m) }
func (*Device_DeviceIdentity) ProtoMessage()               {}
func (*Device_DeviceIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *Device_DeviceIdentity) GetChain() []*common.CertChain {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Device_DeviceIdentity) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type DeviceNetworkTemplate struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description"`
}

func (m *DeviceNetworkTemplate) Reset()                    { *m = DeviceNetworkTemplate{} }
func (m *DeviceNetworkTemplate) String() string            { return proto.CompactTextString(m) }
func (*DeviceNetworkTemplate) ProtoMessage()               {}
func (*DeviceNetworkTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceNetworkTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceNetworkTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceNetworkTemplate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "device.Device")
	proto.RegisterType((*Device_DeviceNetworkSettings)(nil), "device.Device.DeviceNetworkSettings")
	proto.RegisterType((*Device_DeviceInterfaceConfig)(nil), "device.Device.DeviceInterfaceConfig")
	proto.RegisterType((*Device_DeviceInterfaceConfig_WifiConfig)(nil), "device.Device.DeviceInterfaceConfig.WifiConfig")
	proto.RegisterType((*Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork)(nil), "device.Device.DeviceInterfaceConfig.WifiConfig.WifiNetwork")
	proto.RegisterType((*Device_DeviceIdentity)(nil), "device.Device.DeviceIdentity")
	proto.RegisterType((*DeviceNetworkTemplate)(nil), "device.DeviceNetworkTemplate")
	proto.RegisterEnum("device.DeviceStatus", DeviceStatus_name, DeviceStatus_value)
	proto.RegisterEnum("device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType", Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_name, Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_value)
}

func init() { proto.RegisterFile("github.com/fuserobotics/proto/device/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x2e, 0x29, 0xeb, 0xc1, 0x91, 0x1f, 0xea, 0xa2, 0x2d, 0x58, 0xa1, 0x46, 0x55, 0xdb, 0x68,
	0x85, 0x1e, 0xa4, 0xc4, 0x39, 0x05, 0x39, 0x04, 0x8e, 0x6c, 0x20, 0x82, 0x13, 0xc9, 0xa0, 0xed,
	0x38, 0x08, 0x10, 0x08, 0x14, 0x39, 0xa2, 0x17, 0x32, 0x77, 0x37, 0xdc, 0x95, 0x15, 0x5e, 0x72,
	0xcd, 0x21, 0x7f, 0x26, 0xbf, 0x26, 0xbf, 0x27, 0xe0, 0x72, 0xf5, 0xb0, 0xad, 0xbc, 0x90, 0xd3,
	0xce, 0xf3, 0xe3, 0x7c, 0xf3, 0x71, 0xe0, 0x7e, 0x44, 0xd5, 0xe5, 0x64, 0xd8, 0x0a, 0x78, 0xdc,
	0x1e, 0x4d, 0x24, 0x26, 0x7c, 0xc8, 0x15, 0x0d, 0x64, 0x5b, 0x24, 0x5c, 0xf1, 0x76, 0x88, 0xd7,
	0x34, 0x40, 0xf3, 0xb4, 0x74, 0x8c, 0x94, 0x72, 0xaf, 0xfe, 0x8d, 0xd6, 0x80, 0xc7, 0x31, 0x67,
	0xe6, 0xc9, 0x5b, 0x77, 0x3e, 0x3a, 0x50, 0x3a, 0xd4, 0xdd, 0xa4, 0x0e, 0x95, 0x4b, 0x2e, 0x15,
	0xf3, 0x63, 0x74, 0xad, 0x86, 0xd5, 0x74, 0xbc, 0xb9, 0x4f, 0xfe, 0x80, 0x52, 0x82, 0x11, 0xe5,
	0xcc, 0xb5, 0x75, 0xc6, 0x78, 0xa4, 0x0f, 0x35, 0x86, 0x6a, 0xca, 0x93, 0xf1, 0x40, 0xa2, 0x52,
	0x94, 0x45, 0xd2, 0x2d, 0x34, 0xac, 0x66, 0x75, 0x7f, 0xaf, 0x65, 0x46, 0x3c, 0x5c, 0x7e, 0x7a,
	0x79, 0xf1, 0xa9, 0xa9, 0xf5, 0xb6, 0xd8, 0xcd, 0x00, 0x79, 0x08, 0x15, 0x1a, 0x22, 0x53, 0x54,
	0xa5, 0xee, 0x9a, 0x06, 0xda, 0x5e, 0x09, 0xd4, 0x35, 0x45, 0xde, 0xbc, 0xbc, 0xfe, 0x0e, 0x7e,
	0x5f, 0xf9, 0x11, 0xf2, 0x0f, 0xd8, 0x54, 0x68, 0x4a, 0xd5, 0xfd, 0x5f, 0x5b, 0x86, 0x7e, 0xf7,
	0xe4, 0x20, 0x0c, 0x13, 0x94, 0xd2, 0xb3, 0xa9, 0x20, 0x4f, 0xc0, 0xa1, 0x4c, 0x61, 0x32, 0xf2,
	0x03, 0x74, 0xed, 0x46, 0xe1, 0x8b, 0x04, 0xba, 0xb3, 0xaa, 0x0e, 0x67, 0x23, 0x1a, 0x79, 0x8b,
	0xb6, 0xfa, 0x87, 0xd2, 0x6c, 0x80, 0x5b, 0x45, 0xc4, 0x85, 0x72, 0x88, 0xd7, 0x4b, 0x8b, 0x9d,
	0xb9, 0x66, 0x34, 0xfb, 0x6b, 0xa3, 0xdd, 0x03, 0x88, 0x7c, 0x85, 0x53, 0x3f, 0x1d, 0x50, 0x61,
	0x96, 0xbb, 0xa2, 0xd4, 0x31, 0x45, 0x5d, 0x41, 0xce, 0x60, 0x63, 0x4a, 0x47, 0x74, 0xa1, 0x48,
	0xbe, 0xc8, 0xf6, 0xf7, 0x10, 0x6a, 0x5d, 0xd0, 0x11, 0x35, 0xdc, 0xd6, 0x33, 0x94, 0xf9, 0x16,
	0x77, 0xa1, 0x10, 0x32, 0xe9, 0x16, 0xf5, 0x72, 0x56, 0x0c, 0x90, 0x65, 0xeb, 0xef, 0xd7, 0x00,
	0x16, 0x08, 0xe4, 0x1c, 0xca, 0x46, 0x60, 0xd7, 0xd2, 0x7d, 0x8f, 0x7e, 0x70, 0x06, 0x6d, 0x1a,
	0x3d, 0xbd, 0x19, 0x16, 0xd9, 0x85, 0x0d, 0x7c, 0xab, 0x12, 0x7f, 0xc0, 0x85, 0xa2, 0x9c, 0x49,
	0xf3, 0x53, 0xae, 0xeb, 0x60, 0x3f, 0x8f, 0xd5, 0x3f, 0xd9, 0x50, 0x5d, 0xea, 0x26, 0x04, 0xd6,
	0xa4, 0xa4, 0xa1, 0x51, 0x40, 0xdb, 0xe4, 0x15, 0x54, 0x18, 0xaa, 0x81, 0x4a, 0x05, 0xba, 0x4e,
	0xc3, 0x6a, 0x6e, 0xee, 0x3f, 0xfe, 0x89, 0x01, 0xcf, 0x52, 0x81, 0x7a, 0xc8, 0xcc, 0x20, 0x35,
	0x28, 0x08, 0x39, 0x76, 0x4b, 0xfa, 0x73, 0x99, 0x49, 0xfe, 0x86, 0xaa, 0x90, 0xe3, 0x01, 0xb2,
	0x80, 0x87, 0x18, 0xba, 0xd0, 0xb0, 0x9a, 0x15, 0x0f, 0x84, 0x1c, 0x1f, 0xe5, 0x11, 0xf2, 0x1b,
	0x14, 0xf5, 0x55, 0x1a, 0x3e, 0xb9, 0x43, 0xfe, 0x84, 0xca, 0x18, 0xd3, 0x41, 0x1c, 0xc5, 0x4a,
	0xcb, 0xef, 0x78, 0xe5, 0x31, 0xa6, 0xcf, 0xa3, 0x58, 0x65, 0x27, 0x2b, 0x7c, 0x9a, 0x4c, 0xa9,
	0x44, 0x2d, 0xb2, 0xe3, 0xcd, 0xfd, 0x0c, 0x2c, 0x4a, 0xf8, 0x44, 0xb8, 0xc5, 0x1c, 0x4c, 0x3b,
	0xe4, 0x2f, 0x70, 0x46, 0x09, 0xbe, 0x99, 0x20, 0x0b, 0x52, 0xb7, 0xa2, 0x33, 0x8b, 0xc0, 0xdd,
	0xc5, 0x96, 0xef, 0x2e, 0x76, 0x67, 0x0f, 0xb6, 0x6e, 0x91, 0x26, 0x25, 0xb0, 0x0f, 0x4e, 0x6a,
	0xbf, 0x10, 0x07, 0x8a, 0x07, 0x87, 0x4f, 0xfb, 0x9d, 0x9a, 0x55, 0x7f, 0x09, 0x9b, 0x37, 0x2f,
	0x95, 0xfc, 0x07, 0xc5, 0xe0, 0xd2, 0xa7, 0xcc, 0xdc, 0xd7, 0xfc, 0x17, 0xea, 0x60, 0xa2, 0x3a,
	0x59, 0xc2, 0xcb, 0xf3, 0x64, 0x1b, 0x40, 0x4c, 0x86, 0x57, 0x34, 0x18, 0x8c, 0x31, 0x35, 0x94,
	0x9d, 0x3c, 0x72, 0x8c, 0xe9, 0xce, 0xeb, 0x5b, 0x77, 0x7e, 0x86, 0xb1, 0xb8, 0xf2, 0x15, 0x92,
	0x4d, 0xb0, 0xe7, 0xfa, 0xda, 0x34, 0xcc, 0x14, 0xd7, 0x37, 0x97, 0x6f, 0x53, 0xdb, 0xa4, 0x01,
	0xd5, 0x10, 0x65, 0x90, 0x50, 0x4d, 0xc6, 0x80, 0x2f, 0x87, 0xfe, 0xff, 0x17, 0xd6, 0x73, 0xf8,
	0x53, 0xe5, 0xab, 0x89, 0x24, 0x55, 0x28, 0x9f, 0xf7, 0x8e, 0x7b, 0xfd, 0x8b, 0x9e, 0x21, 0xf8,
	0xac, 0xfb, 0xe2, 0xa8, 0x66, 0x0d, 0x4b, 0x5a, 0x9d, 0x07, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0x37, 0x07, 0xac, 0xb0, 0x05, 0x00, 0x00,
}
