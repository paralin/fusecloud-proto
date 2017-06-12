// Code generated by protoc-gen-go.
// source: github.com/fuserobotics/proto/device/device.proto
// DO NOT EDIT!

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	github.com/fuserobotics/proto/device/device.proto

It has these top-level messages:
	Device
	DeviceConnection
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

type DeviceConnectionType int32

const (
	// Ethernet connection
	DeviceConnectionType_CONN_ETH DeviceConnectionType = 0
	// Wifi connection
	DeviceConnectionType_CONN_WIFI DeviceConnectionType = 1
)

var DeviceConnectionType_name = map[int32]string{
	0: "CONN_ETH",
	1: "CONN_WIFI",
}
var DeviceConnectionType_value = map[string]int32{
	"CONN_ETH":  0,
	"CONN_WIFI": 1,
}

func (x DeviceConnectionType) String() string {
	return proto.EnumName(DeviceConnectionType_name, int32(x))
}
func (DeviceConnectionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (DeviceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeviceConnection_WifiConfig_WifiMode int32

const (
	DeviceConnection_WifiConfig_WIFI_INFRA DeviceConnection_WifiConfig_WifiMode = 0
	DeviceConnection_WifiConfig_WIFI_ADHOC DeviceConnection_WifiConfig_WifiMode = 1
	DeviceConnection_WifiConfig_WIFI_AP    DeviceConnection_WifiConfig_WifiMode = 2
)

var DeviceConnection_WifiConfig_WifiMode_name = map[int32]string{
	0: "WIFI_INFRA",
	1: "WIFI_ADHOC",
	2: "WIFI_AP",
}
var DeviceConnection_WifiConfig_WifiMode_value = map[string]int32{
	"WIFI_INFRA": 0,
	"WIFI_ADHOC": 1,
	"WIFI_AP":    2,
}

func (x DeviceConnection_WifiConfig_WifiMode) String() string {
	return proto.EnumName(DeviceConnection_WifiConfig_WifiMode_name, int32(x))
}
func (DeviceConnection_WifiConfig_WifiMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1, 0}
}

type DeviceConnection_WifiSecurity_WifiSecurityMode int32

const (
	// None to explicitly disable security
	DeviceConnection_WifiSecurity_WIFI_SEC_NONE DeviceConnection_WifiSecurity_WifiSecurityMode = 0
	// WPA and WPA2 Personal
	DeviceConnection_WifiSecurity_WIFI_SEC_PERSONAL DeviceConnection_WifiSecurity_WifiSecurityMode = 1
	// WPA and WPA2 Enterprise
	DeviceConnection_WifiSecurity_WIFI_SEC_ENTERPRISE DeviceConnection_WifiSecurity_WifiSecurityMode = 2
)

var DeviceConnection_WifiSecurity_WifiSecurityMode_name = map[int32]string{
	0: "WIFI_SEC_NONE",
	1: "WIFI_SEC_PERSONAL",
	2: "WIFI_SEC_ENTERPRISE",
}
var DeviceConnection_WifiSecurity_WifiSecurityMode_value = map[string]int32{
	"WIFI_SEC_NONE":       0,
	"WIFI_SEC_PERSONAL":   1,
	"WIFI_SEC_ENTERPRISE": 2,
}

func (x DeviceConnection_WifiSecurity_WifiSecurityMode) String() string {
	return proto.EnumName(DeviceConnection_WifiSecurity_WifiSecurityMode_name, int32(x))
}
func (DeviceConnection_WifiSecurity_WifiSecurityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 2, 0}
}

// Device: a full-linux computer running serf, docker stack
type Device struct {
	Hostname        string                        `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Region          string                        `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	NetworkSettings *Device_DeviceNetworkSettings `protobuf:"bytes,3,opt,name=network_settings,json=networkSettings" json:"network_settings,omitempty"`
	Identity        *Device_DeviceIdentity        `protobuf:"bytes,4,opt,name=identity" json:"identity,omitempty"`
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
	Ip *common.IPAddress `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	// reserved field 2, 3
	Connection []*DeviceConnection `protobuf:"bytes,4,rep,name=connection" json:"connection,omitempty"`
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

func (m *Device_DeviceNetworkSettings) GetConnection() []*DeviceConnection {
	if m != nil {
		return m.Connection
	}
	return nil
}

type Device_DeviceIdentity struct {
	// Generated certs, latest is first
	Chain     []*common.CertChain `protobuf:"bytes,2,rep,name=chain" json:"chain,omitempty"`
	PublicKey string              `protobuf:"bytes,3,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
}

func (m *Device_DeviceIdentity) Reset()                    { *m = Device_DeviceIdentity{} }
func (m *Device_DeviceIdentity) String() string            { return proto.CompactTextString(m) }
func (*Device_DeviceIdentity) ProtoMessage()               {}
func (*Device_DeviceIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

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

// DeviceConnectionConfig represents a NetworkManager connection.
type DeviceConnection struct {
	// type controls what kind of connection this is
	Type DeviceConnectionType `protobuf:"varint,1,opt,name=type,enum=device.DeviceConnectionType" json:"type,omitempty"`
	// auto_connect controls if the device should automatically use this connection or not.
	AutoConnect bool `protobuf:"varint,2,opt,name=auto_connect,json=autoConnect" json:"auto_connect,omitempty"`
	// auto_connect_priority: higher numbers -> higher priority
	AutoConnectPriority int32 `protobuf:"varint,3,opt,name=auto_connect_priority,json=autoConnectPriority" json:"auto_connect_priority,omitempty"`
	// auto_connect_retries - how many connection attempts?
	AutoConnectRetries int32 `protobuf:"varint,4,opt,name=auto_connect_retries,json=autoConnectRetries" json:"auto_connect_retries,omitempty"`
	// interface restricts the connection to a specific interface.
	Interface string `protobuf:"bytes,5,opt,name=interface" json:"interface,omitempty"`
	// ipv4 settings
	Ipv4 *DeviceConnection_IPV4Config `protobuf:"bytes,6,opt,name=ipv4" json:"ipv4,omitempty"`
	// wifi settings
	Wifi *DeviceConnection_WifiConfig `protobuf:"bytes,7,opt,name=wifi" json:"wifi,omitempty"`
	// ethernet settings
	Ethernet *DeviceConnection_EthernetConfig `protobuf:"bytes,8,opt,name=ethernet" json:"ethernet,omitempty"`
}

func (m *DeviceConnection) Reset()                    { *m = DeviceConnection{} }
func (m *DeviceConnection) String() string            { return proto.CompactTextString(m) }
func (*DeviceConnection) ProtoMessage()               {}
func (*DeviceConnection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceConnection) GetType() DeviceConnectionType {
	if m != nil {
		return m.Type
	}
	return DeviceConnectionType_CONN_ETH
}

func (m *DeviceConnection) GetAutoConnect() bool {
	if m != nil {
		return m.AutoConnect
	}
	return false
}

func (m *DeviceConnection) GetAutoConnectPriority() int32 {
	if m != nil {
		return m.AutoConnectPriority
	}
	return 0
}

func (m *DeviceConnection) GetAutoConnectRetries() int32 {
	if m != nil {
		return m.AutoConnectRetries
	}
	return 0
}

func (m *DeviceConnection) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *DeviceConnection) GetIpv4() *DeviceConnection_IPV4Config {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *DeviceConnection) GetWifi() *DeviceConnection_WifiConfig {
	if m != nil {
		return m.Wifi
	}
	return nil
}

func (m *DeviceConnection) GetEthernet() *DeviceConnection_EthernetConfig {
	if m != nil {
		return m.Ethernet
	}
	return nil
}

type DeviceConnection_EthernetConfig struct {
}

func (m *DeviceConnection_EthernetConfig) Reset()         { *m = DeviceConnection_EthernetConfig{} }
func (m *DeviceConnection_EthernetConfig) String() string { return proto.CompactTextString(m) }
func (*DeviceConnection_EthernetConfig) ProtoMessage()    {}
func (*DeviceConnection_EthernetConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type DeviceConnection_WifiConfig struct {
	Ssid     string                               `protobuf:"bytes,1,opt,name=ssid" json:"ssid,omitempty"`
	Mode     DeviceConnection_WifiConfig_WifiMode `protobuf:"varint,2,opt,name=mode,enum=device.DeviceConnection_WifiConfig_WifiMode" json:"mode,omitempty"`
	Security *DeviceConnection_WifiSecurity       `protobuf:"bytes,3,opt,name=security" json:"security,omitempty"`
}

func (m *DeviceConnection_WifiConfig) Reset()                    { *m = DeviceConnection_WifiConfig{} }
func (m *DeviceConnection_WifiConfig) String() string            { return proto.CompactTextString(m) }
func (*DeviceConnection_WifiConfig) ProtoMessage()               {}
func (*DeviceConnection_WifiConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *DeviceConnection_WifiConfig) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

func (m *DeviceConnection_WifiConfig) GetMode() DeviceConnection_WifiConfig_WifiMode {
	if m != nil {
		return m.Mode
	}
	return DeviceConnection_WifiConfig_WIFI_INFRA
}

func (m *DeviceConnection_WifiConfig) GetSecurity() *DeviceConnection_WifiSecurity {
	if m != nil {
		return m.Security
	}
	return nil
}

type DeviceConnection_WifiSecurity struct {
	Mode DeviceConnection_WifiSecurity_WifiSecurityMode `protobuf:"varint,1,opt,name=mode,enum=device.DeviceConnection_WifiSecurity_WifiSecurityMode" json:"mode,omitempty"`
	// username, if using enterprise.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// password, if using personal
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *DeviceConnection_WifiSecurity) Reset()         { *m = DeviceConnection_WifiSecurity{} }
func (m *DeviceConnection_WifiSecurity) String() string { return proto.CompactTextString(m) }
func (*DeviceConnection_WifiSecurity) ProtoMessage()    {}
func (*DeviceConnection_WifiSecurity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 2}
}

func (m *DeviceConnection_WifiSecurity) GetMode() DeviceConnection_WifiSecurity_WifiSecurityMode {
	if m != nil {
		return m.Mode
	}
	return DeviceConnection_WifiSecurity_WIFI_SEC_NONE
}

func (m *DeviceConnection_WifiSecurity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DeviceConnection_WifiSecurity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeviceConnection_IPV4Config struct {
	// If set, will use manual mode instead of dhcp.
	Address   []*DeviceConnection_IPV4ConfigAddress `protobuf:"bytes,1,rep,name=address" json:"address,omitempty"`
	Dns       []*common.IPAddress                   `protobuf:"bytes,2,rep,name=dns" json:"dns,omitempty"`
	DnsSearch []string                              `protobuf:"bytes,3,rep,name=dns_search,json=dnsSearch" json:"dns_search,omitempty"`
	// if set, we will ignore DNS servers given by the DHCP server
	DisableAutoDns bool `protobuf:"varint,4,opt,name=disable_auto_dns,json=disableAutoDns" json:"disable_auto_dns,omitempty"`
}

func (m *DeviceConnection_IPV4Config) Reset()                    { *m = DeviceConnection_IPV4Config{} }
func (m *DeviceConnection_IPV4Config) String() string            { return proto.CompactTextString(m) }
func (*DeviceConnection_IPV4Config) ProtoMessage()               {}
func (*DeviceConnection_IPV4Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

func (m *DeviceConnection_IPV4Config) GetAddress() []*DeviceConnection_IPV4ConfigAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DeviceConnection_IPV4Config) GetDns() []*common.IPAddress {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *DeviceConnection_IPV4Config) GetDnsSearch() []string {
	if m != nil {
		return m.DnsSearch
	}
	return nil
}

func (m *DeviceConnection_IPV4Config) GetDisableAutoDns() bool {
	if m != nil {
		return m.DisableAutoDns
	}
	return false
}

// Configured address for an IPV4 device.
type DeviceConnection_IPV4ConfigAddress struct {
	Range   *common.IPRange   `protobuf:"bytes,1,opt,name=range" json:"range,omitempty"`
	Gateway *common.IPAddress `protobuf:"bytes,2,opt,name=gateway" json:"gateway,omitempty"`
}

func (m *DeviceConnection_IPV4ConfigAddress) Reset()         { *m = DeviceConnection_IPV4ConfigAddress{} }
func (m *DeviceConnection_IPV4ConfigAddress) String() string { return proto.CompactTextString(m) }
func (*DeviceConnection_IPV4ConfigAddress) ProtoMessage()    {}
func (*DeviceConnection_IPV4ConfigAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 4}
}

func (m *DeviceConnection_IPV4ConfigAddress) GetRange() *common.IPRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *DeviceConnection_IPV4ConfigAddress) GetGateway() *common.IPAddress {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "device.Device")
	proto.RegisterType((*Device_DeviceNetworkSettings)(nil), "device.Device.DeviceNetworkSettings")
	proto.RegisterType((*Device_DeviceIdentity)(nil), "device.Device.DeviceIdentity")
	proto.RegisterType((*DeviceConnection)(nil), "device.DeviceConnection")
	proto.RegisterType((*DeviceConnection_EthernetConfig)(nil), "device.DeviceConnection.EthernetConfig")
	proto.RegisterType((*DeviceConnection_WifiConfig)(nil), "device.DeviceConnection.WifiConfig")
	proto.RegisterType((*DeviceConnection_WifiSecurity)(nil), "device.DeviceConnection.WifiSecurity")
	proto.RegisterType((*DeviceConnection_IPV4Config)(nil), "device.DeviceConnection.IPV4Config")
	proto.RegisterType((*DeviceConnection_IPV4ConfigAddress)(nil), "device.DeviceConnection.IPV4ConfigAddress")
	proto.RegisterEnum("device.DeviceConnectionType", DeviceConnectionType_name, DeviceConnectionType_value)
	proto.RegisterEnum("device.DeviceStatus", DeviceStatus_name, DeviceStatus_value)
	proto.RegisterEnum("device.DeviceConnection_WifiConfig_WifiMode", DeviceConnection_WifiConfig_WifiMode_name, DeviceConnection_WifiConfig_WifiMode_value)
	proto.RegisterEnum("device.DeviceConnection_WifiSecurity_WifiSecurityMode", DeviceConnection_WifiSecurity_WifiSecurityMode_name, DeviceConnection_WifiSecurity_WifiSecurityMode_value)
}

func init() { proto.RegisterFile("github.com/fuserobotics/proto/device/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 842 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0xf5, 0x67, 0x69, 0xe4, 0x28, 0xf4, 0x24, 0x6e, 0x09, 0x22, 0x01, 0x1c, 0xa7, 0x69,
	0x0c, 0xb7, 0x90, 0x53, 0x25, 0x68, 0xda, 0xb7, 0x0a, 0x12, 0x83, 0xa8, 0x49, 0x29, 0x61, 0xe5,
	0xc4, 0x7d, 0x23, 0x28, 0x72, 0x25, 0x2d, 0x12, 0x71, 0x09, 0xee, 0x2a, 0x86, 0x6e, 0xd2, 0x4b,
	0xf4, 0x0e, 0x3d, 0x49, 0xdf, 0x7b, 0x8b, 0x62, 0x97, 0x2b, 0x5a, 0x72, 0xed, 0xda, 0x4f, 0xd4,
	0xcc, 0x7c, 0xdf, 0xec, 0x37, 0x3b, 0x3b, 0x23, 0xf8, 0x61, 0xc6, 0xe4, 0x7c, 0x39, 0x69, 0x47,
	0x7c, 0x71, 0x3a, 0x5d, 0x0a, 0x9a, 0xf1, 0x09, 0x97, 0x2c, 0x12, 0xa7, 0x69, 0xc6, 0x25, 0x3f,
	0x8d, 0xe9, 0x17, 0x16, 0x51, 0xf3, 0x69, 0x6b, 0x1f, 0xd6, 0x72, 0xcb, 0xbd, 0x85, 0x1a, 0xf1,
	0xc5, 0x82, 0x27, 0xe6, 0x93, 0x53, 0x8f, 0xfe, 0x28, 0x43, 0xad, 0xaf, 0xd9, 0xe8, 0x42, 0x7d,
	0xce, 0x85, 0x4c, 0xc2, 0x05, 0x75, 0xac, 0x43, 0xeb, 0xb8, 0x41, 0x0a, 0x1b, 0xbf, 0x82, 0x5a,
	0x46, 0x67, 0x8c, 0x27, 0x4e, 0x49, 0x47, 0x8c, 0x85, 0x43, 0xb0, 0x13, 0x2a, 0x2f, 0x78, 0xf6,
	0x29, 0x10, 0x54, 0x4a, 0x96, 0xcc, 0x84, 0x53, 0x3e, 0xb4, 0x8e, 0x9b, 0x9d, 0x6f, 0xda, 0x46,
	0x62, 0x7f, 0xf3, 0xe3, 0xe7, 0xe0, 0xb1, 0xc1, 0x92, 0xfb, 0xc9, 0xb6, 0x03, 0x7f, 0x86, 0x3a,
	0x8b, 0x69, 0x22, 0x99, 0x5c, 0x39, 0x15, 0x9d, 0xe8, 0xf1, 0xb5, 0x89, 0x06, 0x06, 0x44, 0x0a,
	0xb8, 0x2b, 0xe1, 0xe0, 0xda, 0x43, 0xf0, 0x09, 0x94, 0x58, 0xaa, 0x4b, 0x6a, 0x76, 0xf6, 0xdb,
	0xa6, 0xfc, 0xc1, 0xa8, 0x1b, 0xc7, 0x19, 0x15, 0x82, 0x94, 0x58, 0x8a, 0x3f, 0x01, 0x44, 0x3c,
	0x49, 0x68, 0x24, 0x55, 0x8d, 0x95, 0xc3, 0xf2, 0x71, 0xb3, 0xe3, 0x6c, 0x1f, 0xdc, 0x2b, 0xe2,
	0x64, 0x03, 0xeb, 0xfe, 0x0e, 0xad, 0x6d, 0x45, 0xf8, 0x1c, 0xaa, 0xd1, 0x3c, 0x64, 0xea, 0xaa,
	0xca, 0x9b, 0x27, 0xf6, 0x68, 0x26, 0x7b, 0x2a, 0x40, 0xf2, 0x38, 0x3e, 0x06, 0x48, 0x97, 0x93,
	0xcf, 0x2c, 0x0a, 0x3e, 0xd1, 0x95, 0xbe, 0xb6, 0x06, 0x69, 0xe4, 0x9e, 0x77, 0x74, 0x75, 0xf4,
	0x67, 0x03, 0xec, 0xab, 0x47, 0xe3, 0x0b, 0xa8, 0xc8, 0x55, 0x9a, 0x37, 0xa8, 0xd5, 0x79, 0x74,
	0x93, 0xc4, 0xb3, 0x55, 0x4a, 0x89, 0x46, 0xe2, 0x13, 0xd8, 0x0b, 0x97, 0x92, 0x07, 0x46, 0xb3,
	0x6e, 0x60, 0x9d, 0x34, 0x95, 0xcf, 0xe0, 0xb1, 0x03, 0x07, 0x9b, 0x90, 0x20, 0xcd, 0x18, 0xcf,
	0x54, 0x07, 0x94, 0xa6, 0x2a, 0x79, 0xb0, 0x81, 0x1d, 0x99, 0x10, 0xbe, 0x80, 0x87, 0x5b, 0x9c,
	0x8c, 0xca, 0x8c, 0x51, 0xa1, 0x9b, 0x56, 0x25, 0xb8, 0x41, 0x21, 0x79, 0x04, 0x1f, 0x41, 0x83,
	0x25, 0x92, 0x66, 0xd3, 0x30, 0xa2, 0x4e, 0x35, 0xaf, 0xb6, 0x70, 0xe0, 0x6b, 0xa8, 0xb0, 0xf4,
	0xcb, 0x2b, 0xa7, 0xa6, 0xdb, 0xf4, 0xf4, 0xa6, 0xc2, 0xda, 0x83, 0xd1, 0xc7, 0x57, 0x3d, 0x9e,
	0x4c, 0xd9, 0x8c, 0x68, 0x82, 0x22, 0x5e, 0xb0, 0x29, 0x73, 0x76, 0x6f, 0x21, 0x9e, 0xb3, 0x29,
	0x5b, 0x13, 0x15, 0x01, 0x7b, 0x50, 0xa7, 0x72, 0x4e, 0xb3, 0x84, 0x4a, 0xa7, 0xae, 0xc9, 0xcf,
	0x6f, 0x24, 0x7b, 0x06, 0x68, 0x12, 0x14, 0x44, 0xd7, 0x86, 0xd6, 0x76, 0xcc, 0xfd, 0xdb, 0x02,
	0xb8, 0x3c, 0x0b, 0x11, 0x2a, 0x42, 0xb0, 0xd8, 0x4c, 0x94, 0xfe, 0x8d, 0xbf, 0x40, 0x65, 0xc1,
	0x63, 0xaa, 0x5b, 0xd1, 0xea, 0x7c, 0x7f, 0x07, 0xc9, 0xfa, 0xe7, 0x6f, 0x3c, 0xa6, 0x44, 0x33,
	0xb1, 0x0b, 0x75, 0x41, 0xa3, 0x65, 0xd1, 0xa4, 0x66, 0xe7, 0xd9, 0xff, 0x66, 0x19, 0x1b, 0x30,
	0x29, 0x68, 0x47, 0xaf, 0xa1, 0xbe, 0x4e, 0x8a, 0x2d, 0x80, 0xf3, 0xc1, 0x9b, 0x41, 0x30, 0xf0,
	0xdf, 0x90, 0xae, 0xbd, 0x53, 0xd8, 0xdd, 0xfe, 0xdb, 0x61, 0xcf, 0xb6, 0xb0, 0x09, 0xbb, 0xb9,
	0x3d, 0xb2, 0x4b, 0xee, 0x3f, 0x16, 0xec, 0x6d, 0xe6, 0xc4, 0x5f, 0x4d, 0x39, 0xf9, 0x9b, 0xfc,
	0xf1, 0x4e, 0x42, 0xb6, 0x8c, 0x8d, 0xc2, 0x5c, 0xa8, 0xab, 0xc5, 0xa5, 0x97, 0x50, 0xbe, 0x6a,
	0x0a, 0x5b, 0xc5, 0xd2, 0x50, 0x88, 0x0b, 0x9e, 0xc5, 0x66, 0x5a, 0x0a, 0xfb, 0xe8, 0x03, 0xd8,
	0x57, 0x33, 0xe2, 0x3e, 0xdc, 0xd3, 0xaa, 0xc7, 0x5e, 0x2f, 0xf0, 0x87, 0xbe, 0x67, 0xef, 0xe0,
	0x01, 0xec, 0x17, 0xae, 0x91, 0x47, 0xc6, 0x43, 0xbf, 0xfb, 0xde, 0xb6, 0xf0, 0x6b, 0x78, 0x50,
	0xb8, 0x3d, 0xff, 0xcc, 0x23, 0x23, 0x32, 0x18, 0x7b, 0x76, 0xc9, 0xfd, 0xcb, 0x02, 0xb8, 0x7c,
	0x71, 0xd8, 0x87, 0xdd, 0x30, 0xdf, 0x1a, 0x8e, 0xa5, 0x87, 0xfb, 0xe4, 0x0e, 0xef, 0x74, 0xbd,
	0x67, 0xd6, 0x54, 0x7c, 0x0a, 0xe5, 0x38, 0x11, 0x57, 0xd7, 0xc3, 0xe5, 0x42, 0x52, 0x51, 0xb5,
	0x1c, 0xe2, 0x44, 0x04, 0x82, 0x86, 0x59, 0x34, 0x77, 0xca, 0x87, 0x65, 0x35, 0x2e, 0x71, 0x22,
	0xc6, 0xda, 0x81, 0xc7, 0x60, 0xc7, 0x4c, 0x84, 0x93, 0xcf, 0x34, 0xd0, 0x63, 0xa8, 0x12, 0x56,
	0xf4, 0x64, 0xb7, 0x8c, 0xbf, 0xbb, 0x94, 0xbc, 0x9f, 0x08, 0x77, 0x06, 0xfb, 0xff, 0xd1, 0x82,
	0xcf, 0xa0, 0x9a, 0x85, 0xc9, 0x8c, 0x9a, 0xad, 0x78, 0xff, 0x52, 0x04, 0x51, 0x6e, 0x92, 0x47,
	0xf1, 0x3b, 0xd8, 0x9d, 0x85, 0x92, 0x5e, 0x84, 0x2b, 0xdd, 0x8c, 0x6b, 0xd5, 0xae, 0x11, 0x27,
	0x2f, 0xe1, 0xe1, 0x75, 0x6b, 0x08, 0xf7, 0xa0, 0xde, 0x1b, 0xfa, 0x7e, 0xe0, 0x9d, 0xbd, 0xb5,
	0x77, 0xf0, 0x1e, 0x34, 0xb4, 0xa5, 0xee, 0xdb, 0xb6, 0x4e, 0xbe, 0x85, 0xbd, 0x9c, 0x34, 0x96,
	0xa1, 0x5c, 0x0a, 0xf5, 0xd2, 0x3e, 0xf8, 0xef, 0xfc, 0xe1, 0xb9, 0x6f, 0xef, 0x60, 0x03, 0xaa,
	0xdd, 0xf7, 0x83, 0x8f, 0x9e, 0x6d, 0x4d, 0x6a, 0xfa, 0xef, 0xea, 0xe5, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x76, 0xfe, 0xb2, 0x8b, 0x1e, 0x07, 0x00, 0x00,
}
