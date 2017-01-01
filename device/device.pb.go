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
	proto.RegisterEnum("device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType", Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_name, Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkType_value)
}

func init() { proto.RegisterFile("github.com/fuserobotics/proto/device/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x25, 0x4e, 0xf3, 0xf0, 0x4d, 0x1f, 0x61, 0x04, 0xc8, 0x58, 0x54, 0x84, 0xb6, 0x12, 0x59,
	0x25, 0x50, 0x56, 0x88, 0x05, 0x2a, 0x29, 0x12, 0x11, 0x82, 0x56, 0xa6, 0x08, 0x84, 0x84, 0x2c,
	0xc7, 0xbe, 0x71, 0x47, 0xae, 0x67, 0x06, 0xcf, 0xa4, 0xc1, 0x1b, 0xb6, 0x2c, 0xf8, 0x19, 0xbe,
	0x86, 0xef, 0x41, 0x1e, 0x4f, 0x1e, 0x6d, 0xc3, 0x4b, 0xac, 0xe6, 0xbe, 0x7d, 0xcf, 0x39, 0xbe,
	0xf0, 0x30, 0xa6, 0xea, 0x74, 0x32, 0xea, 0x85, 0x3c, 0xed, 0x8f, 0x27, 0x12, 0x33, 0x3e, 0xe2,
	0x8a, 0x86, 0xb2, 0x2f, 0x32, 0xae, 0x78, 0x3f, 0xc2, 0x73, 0x1a, 0xa2, 0x79, 0x7a, 0x3a, 0x46,
	0xea, 0xa5, 0xe7, 0xfe, 0xa1, 0x35, 0xe4, 0x69, 0xca, 0x99, 0x79, 0xca, 0xd6, 0x9d, 0xef, 0x36,
	0xd4, 0x0f, 0x75, 0x37, 0x71, 0xa1, 0x79, 0xca, 0xa5, 0x62, 0x41, 0x8a, 0x4e, 0xa5, 0x53, 0xe9,
	0xda, 0xde, 0xdc, 0x27, 0xb7, 0xa0, 0x9e, 0x61, 0x4c, 0x39, 0x73, 0x2c, 0x9d, 0x31, 0x1e, 0x39,
	0x82, 0x36, 0x43, 0x35, 0xe5, 0x59, 0xe2, 0x4b, 0x54, 0x8a, 0xb2, 0x58, 0x3a, 0xd5, 0x4e, 0xa5,
	0xdb, 0xda, 0xdf, 0xeb, 0x99, 0x15, 0x0f, 0x97, 0x9f, 0xd7, 0x65, 0xf1, 0x1b, 0x53, 0xeb, 0x6d,
	0xb1, 0x8b, 0x01, 0xf2, 0x18, 0x9a, 0x34, 0x42, 0xa6, 0xa8, 0xca, 0x9d, 0x35, 0x3d, 0x68, 0x7b,
	0xe5, 0xa0, 0xa1, 0x29, 0xf2, 0xe6, 0xe5, 0xee, 0x17, 0xb8, 0xb9, 0xf2, 0x23, 0xe4, 0x1e, 0x58,
	0x54, 0x68, 0x48, 0xad, 0xfd, 0xeb, 0x3d, 0x03, 0x7f, 0x78, 0x7c, 0x10, 0x45, 0x19, 0x4a, 0xe9,
	0x59, 0x54, 0x90, 0x67, 0x60, 0x53, 0xa6, 0x30, 0x1b, 0x07, 0x21, 0x3a, 0x56, 0xa7, 0xfa, 0x4b,
	0x00, 0xc3, 0x59, 0xd5, 0x80, 0xb3, 0x31, 0x8d, 0xbd, 0x45, 0x9b, 0xfb, 0xad, 0x3e, 0x5b, 0xe0,
	0x52, 0x11, 0x71, 0xa0, 0x11, 0xe1, 0xf9, 0x12, 0xb1, 0x33, 0xd7, 0xac, 0x66, 0xfd, 0x6e, 0xb5,
	0x07, 0x00, 0x71, 0xa0, 0x70, 0x1a, 0xe4, 0x3e, 0x15, 0x86, 0xdc, 0x15, 0xa5, 0xb6, 0x29, 0x1a,
	0x0a, 0x72, 0x02, 0x1b, 0x53, 0x3a, 0xa6, 0x0b, 0x45, 0x4a, 0x22, 0xfb, 0x7f, 0x03, 0xa8, 0xf7,
	0x8e, 0x8e, 0xa9, 0xc1, 0xb6, 0x5e, 0x4c, 0x99, 0xb3, 0xb8, 0x0b, 0xd5, 0x88, 0x49, 0xa7, 0xa6,
	0xc9, 0x59, 0xb1, 0x40, 0x91, 0x75, 0xbf, 0xae, 0x01, 0x2c, 0x26, 0x90, 0xb7, 0xd0, 0x30, 0x02,
	0x3b, 0x15, 0xdd, 0xf7, 0xe4, 0x1f, 0x77, 0xd0, 0xa6, 0xd1, 0xd3, 0x9b, 0xcd, 0x22, 0xbb, 0xb0,
	0x81, 0x9f, 0x55, 0x16, 0xf8, 0x5c, 0x28, 0xca, 0x99, 0x34, 0x3f, 0xe5, 0xba, 0x0e, 0x1e, 0x95,
	0x31, 0xf7, 0x87, 0x05, 0xad, 0xa5, 0x6e, 0x42, 0x60, 0x4d, 0x4a, 0x1a, 0x19, 0x05, 0xb4, 0x4d,
	0x3e, 0x40, 0x93, 0xa1, 0xf2, 0x55, 0x2e, 0xd0, 0xb1, 0x3b, 0x95, 0xee, 0xe6, 0xfe, 0xd3, 0xff,
	0x58, 0xf0, 0x24, 0x17, 0xa8, 0x97, 0x2c, 0x0c, 0xd2, 0x86, 0xaa, 0x90, 0x89, 0x53, 0xd7, 0x9f,
	0x2b, 0x4c, 0x72, 0x17, 0x5a, 0x42, 0x26, 0x3e, 0xb2, 0x90, 0x47, 0x18, 0x39, 0xd0, 0xa9, 0x74,
	0x9b, 0x1e, 0x08, 0x99, 0x3c, 0x2f, 0x23, 0xe4, 0x06, 0xd4, 0xf4, 0x55, 0x1a, 0x3c, 0xa5, 0x43,
	0x6e, 0x43, 0x33, 0xc1, 0xdc, 0x4f, 0xe3, 0x54, 0x69, 0xf9, 0x6d, 0xaf, 0x91, 0x60, 0xfe, 0x2a,
	0x4e, 0x55, 0x71, 0xb2, 0x22, 0xa0, 0xd9, 0x94, 0x4a, 0xd4, 0x22, 0xdb, 0xde, 0xdc, 0x2f, 0x86,
	0xc5, 0x19, 0x9f, 0x08, 0xa7, 0x56, 0x0e, 0xd3, 0x0e, 0xb9, 0x03, 0xf6, 0x38, 0xc3, 0x4f, 0x13,
	0x64, 0x61, 0xee, 0x34, 0x75, 0x66, 0x11, 0xb8, 0x4a, 0x6c, 0xe3, 0x2a, 0xb1, 0x3b, 0x7b, 0xb0,
	0x75, 0x09, 0x34, 0xa9, 0x83, 0x75, 0x70, 0xdc, 0xbe, 0x46, 0x6c, 0xa8, 0x1d, 0x1c, 0xbe, 0x38,
	0x1a, 0xb4, 0x2b, 0xee, 0x7b, 0xd8, 0xbc, 0x78, 0xa9, 0xe4, 0x3e, 0xd4, 0xc2, 0xd3, 0x80, 0x32,
	0x73, 0x5f, 0xf3, 0x5f, 0x68, 0x80, 0x99, 0x1a, 0x14, 0x09, 0xaf, 0xcc, 0x93, 0x6d, 0x00, 0x31,
	0x19, 0x9d, 0xd1, 0xd0, 0x4f, 0x30, 0x37, 0x90, 0xed, 0x32, 0xf2, 0x12, 0xf3, 0x9d, 0x8f, 0x97,
	0xee, 0xfc, 0x04, 0x53, 0x71, 0x16, 0x28, 0x24, 0x9b, 0x60, 0xcd, 0xf5, 0xb5, 0x68, 0x54, 0x28,
	0xae, 0x6f, 0xae, 0x64, 0x53, 0xdb, 0xa4, 0x03, 0xad, 0x08, 0x65, 0x98, 0x51, 0x0d, 0xc6, 0x0c,
	0x5f, 0x0e, 0x8d, 0xea, 0x9a, 0xf5, 0x47, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x93, 0x02, 0x19,
	0x57, 0x88, 0x05, 0x00, 0x00,
}
