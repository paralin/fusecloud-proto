package proto

// Map of GraphQL simplified names to proto names.
var GraphQLTypes = map[string]string{
	"Device":                "github.com/fuserobotics/proto/device.Device",
	"DeviceIdentity":        "github.com/fuserobotics/proto/device.Device_DeviceIdentity",
	"DeviceNetworkTemplate": "github.com/fuserobotics/proto/device.DeviceNetworkTemplate",
	"DeviceInterfaceConfig": "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig",
	"WifiConfig":            "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig",
	"KVGDeviceSubKeys":      "github.com/fuserobotics/proto/device.KVGDeviceSubKeys",
	"UserPrivateData":       "github.com/fuserobotics/proto/auth.UserPrivateData",
	"ScryptParams":          "github.com/fuserobotics/proto/auth.UserPrivateData_ScryptParams",
	"CertChain":             "github.com/fuserobotics/proto/common.CertChain",
	"Region":                "github.com/fuserobotics/proto/region.Region",
	"IPAddress":             "github.com/fuserobotics/proto/common.IPAddress",
	"DeviceStatus":          "github.com/fuserobotics/proto/device.DeviceStatus",
	"WifiNetwork":           "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork",
	"DeviceNetworkSettings": "github.com/fuserobotics/proto/device.Device_DeviceNetworkSettings",
	"GetUnixTimeResponse":   "github.com/fuserobotics/proto/timesync.GetUnixTimeResponse",
	"UserPublicIdentity":    "github.com/fuserobotics/proto/auth.UserPublicIdentity",
	"GeoLocation":           "github.com/fuserobotics/proto/common.GeoLocation",
	"IPRange":               "github.com/fuserobotics/proto/common.IPRange",
}
