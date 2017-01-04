// +build !codegen

package resolvers

import (
	"github.com/fuserobotics/proto/auth"
	"github.com/fuserobotics/proto/common"
	"github.com/fuserobotics/proto/device"
	"github.com/fuserobotics/proto/region"
	"github.com/fuserobotics/proto/timesync"
)

type UserPrivateDataResolver struct {
	Base *auth.UserPrivateData
}

func (r *UserPrivateDataResolver) PrivateKey() *[]*uint8 {
	res := make([]*uint8, len(r.Base.PrivateKey))
	for idx, val := range r.Base.PrivateKey {
		res[idx] = &val
	}
	return &res
}

func (r *UserPrivateDataResolver) Strategy() *int32 {
	res := (*int32)(&r.Base.Strategy)
	return res
}

func (r *UserPrivateDataResolver) ScryptParams() *UserPrivateData_ScryptParamsResolver {
	if r.Base.ScryptParams == nil {
		return nil
	}
	return &UserPrivateData_ScryptParamsResolver{Base: r.Base.ScryptParams}
}

type UserPrivateData_ScryptParamsResolver struct {
	Base *auth.UserPrivateData_ScryptParams
}

func (r *UserPrivateData_ScryptParamsResolver) N() *uint32 {
	return &r.Base.N
}

func (r *UserPrivateData_ScryptParamsResolver) R() *uint32 {
	return &r.Base.R
}

func (r *UserPrivateData_ScryptParamsResolver) P() *uint32 {
	return &r.Base.P
}

type UserPublicIdentityResolver struct {
	Base *auth.UserPublicIdentity
}

func (r *UserPublicIdentityResolver) Username() *string {
	return &r.Base.Username
}

func (r *UserPublicIdentityResolver) PublicKey() *string {
	return &r.Base.PublicKey
}

type CertChainResolver struct {
	Base *common.CertChain
}

func (r *CertChainResolver) Cert() *[]*string {
	res := make([]*string, len(r.Base.Cert))
	for idx, val := range r.Base.Cert {
		res[idx] = &val
	}
	return &res
}

func (r *CertChainResolver) ValidUntil() *int64 {
	return &r.Base.ValidUntil
}

func (r *CertChainResolver) Csr() *string {
	return &r.Base.Csr
}

func (r *CertChainResolver) CsrWaiting() *bool {
	return &r.Base.CsrWaiting
}

type GeoLocationResolver struct {
	Base *common.GeoLocation
}

func (r *GeoLocationResolver) Latitude() *float64 {
	return &r.Base.Latitude
}

func (r *GeoLocationResolver) Longitude() *float64 {
	return &r.Base.Longitude
}

type IPAddressResolver struct {
	Base *common.IPAddress
}

func (r *IPAddressResolver) Address() *[]*uint32 {
	res := make([]*uint32, len(r.Base.Address))
	for idx, val := range r.Base.Address {
		res[idx] = &val
	}
	return &res
}

type IPRangeResolver struct {
	Base *common.IPRange
}

func (r *IPRangeResolver) Ip() *IPAddressResolver {
	if r.Base.Ip == nil {
		return nil
	}
	return &IPAddressResolver{Base: r.Base.Ip}
}

func (r *IPRangeResolver) Plen() *uint32 {
	return &r.Base.Plen
}

type DeviceResolver struct {
	Base *device.Device
}

func (r *DeviceResolver) Hostname() *string {
	return &r.Base.Hostname
}

func (r *DeviceResolver) Region() *string {
	return &r.Base.Region
}

func (r *DeviceResolver) NetworkSettings() *Device_DeviceNetworkSettingsResolver {
	if r.Base.NetworkSettings == nil {
		return nil
	}
	return &Device_DeviceNetworkSettingsResolver{Base: r.Base.NetworkSettings}
}

func (r *DeviceResolver) Identity() *Device_DeviceIdentityResolver {
	if r.Base.Identity == nil {
		return nil
	}
	return &Device_DeviceIdentityResolver{Base: r.Base.Identity}
}

type DeviceNetworkTemplateResolver struct {
	Base *device.DeviceNetworkTemplate
}

func (r *DeviceNetworkTemplateResolver) Id() *string {
	return &r.Base.Id
}

func (r *DeviceNetworkTemplateResolver) Name() *string {
	return &r.Base.Name
}

func (r *DeviceNetworkTemplateResolver) Description() *string {
	return &r.Base.Description
}

type DeviceWithStatusResolver struct {
	Base *device.DeviceWithStatus
}

func (r *DeviceWithStatusResolver) Device() *DeviceResolver {
	if r.Base.Device == nil {
		return nil
	}
	return &DeviceResolver{Base: r.Base.Device}
}

func (r *DeviceWithStatusResolver) Status() *int32 {
	res := (*int32)(&r.Base.Status)
	return res
}

type Device_DeviceIdentityResolver struct {
	Base *device.Device_DeviceIdentity
}

func (r *Device_DeviceIdentityResolver) Chain() *[]*CertChainResolver {
	res := make([]*CertChainResolver, len(r.Base.Chain))
	for idx, val := range r.Base.Chain {
		res[idx] = &CertChainResolver{Base: val}
	}
	return &res
}

func (r *Device_DeviceIdentityResolver) PublicKey() *string {
	return &r.Base.PublicKey
}

type Device_DeviceInterfaceConfigResolver struct {
	Base *device.Device_DeviceInterfaceConfig
}

func (r *Device_DeviceInterfaceConfigResolver) Devname() *string {
	return &r.Base.Devname
}

func (r *Device_DeviceInterfaceConfigResolver) Ip() *IPAddressResolver {
	if r.Base.Ip == nil {
		return nil
	}
	return &IPAddressResolver{Base: r.Base.Ip}
}

func (r *Device_DeviceInterfaceConfigResolver) GatewayIp() *IPAddressResolver {
	if r.Base.GatewayIp == nil {
		return nil
	}
	return &IPAddressResolver{Base: r.Base.GatewayIp}
}

func (r *Device_DeviceInterfaceConfigResolver) WifiSettings() *Device_DeviceInterfaceConfig_WifiConfigResolver {
	if r.Base.WifiSettings == nil {
		return nil
	}
	return &Device_DeviceInterfaceConfig_WifiConfigResolver{Base: r.Base.WifiSettings}
}

func (r *Device_DeviceInterfaceConfigResolver) Dns() *[]*IPAddressResolver {
	res := make([]*IPAddressResolver, len(r.Base.Dns))
	for idx, val := range r.Base.Dns {
		res[idx] = &IPAddressResolver{Base: val}
	}
	return &res
}

type Device_DeviceInterfaceConfig_WifiConfigResolver struct {
	Base *device.Device_DeviceInterfaceConfig_WifiConfig
}

func (r *Device_DeviceInterfaceConfig_WifiConfigResolver) Network() *[]*Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver {
	res := make([]*Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver, len(r.Base.Network))
	for idx, val := range r.Base.Network {
		res[idx] = &Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver{Base: val}
	}
	return &res
}

func (r *Device_DeviceInterfaceConfig_WifiConfigResolver) ExtraOptions() *string {
	return &r.Base.ExtraOptions
}

type Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver struct {
	Base *device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Ssid() *string {
	return &r.Base.Ssid
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) NetType() *int32 {
	res := (*int32)(&r.Base.NetType)
	return res
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Psk() *string {
	return &r.Base.Psk
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) PskEncoded() *bool {
	return &r.Base.PskEncoded
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Proto() *string {
	return &r.Base.Proto
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) KeyMgmt() *string {
	return &r.Base.KeyMgmt
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Pairwise() *string {
	return &r.Base.Pairwise
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Group() *string {
	return &r.Base.Group
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) Frequency() *string {
	return &r.Base.Frequency
}

func (r *Device_DeviceInterfaceConfig_WifiConfig_WifiNetworkResolver) ExtraOptions() *string {
	return &r.Base.ExtraOptions
}

type Device_DeviceNetworkSettingsResolver struct {
	Base *device.Device_DeviceNetworkSettings
}

func (r *Device_DeviceNetworkSettingsResolver) Ip() *IPAddressResolver {
	if r.Base.Ip == nil {
		return nil
	}
	return &IPAddressResolver{Base: r.Base.Ip}
}

func (r *Device_DeviceNetworkSettingsResolver) Interface() *[]*Device_DeviceInterfaceConfigResolver {
	res := make([]*Device_DeviceInterfaceConfigResolver, len(r.Base.Interface))
	for idx, val := range r.Base.Interface {
		res[idx] = &Device_DeviceInterfaceConfigResolver{Base: val}
	}
	return &res
}

type KVGDeviceSubKeysResolver struct {
	Base *device.KVGDeviceSubKeys
}

func (r *KVGDeviceSubKeysResolver) DeviceInfo() *string {
	return &r.Base.DeviceInfo
}

type RegionResolver struct {
	Base *region.Region
}

func (r *RegionResolver) Id() *string {
	return &r.Base.Id
}

func (r *RegionResolver) Name() *string {
	return &r.Base.Name
}

func (r *RegionResolver) IpRange() *IPRangeResolver {
	if r.Base.IpRange == nil {
		return nil
	}
	return &IPRangeResolver{Base: r.Base.IpRange}
}

func (r *RegionResolver) Location() *GeoLocationResolver {
	if r.Base.Location == nil {
		return nil
	}
	return &GeoLocationResolver{Base: r.Base.Location}
}

func (r *RegionResolver) Zoomlevel() *int32 {
	return &r.Base.Zoomlevel
}

type GetUnixTimeResponseResolver struct {
	Base *timesync.GetUnixTimeResponse
}

func (r *GetUnixTimeResponseResolver) Time() *uint32 {
	return &r.Base.Time
}
