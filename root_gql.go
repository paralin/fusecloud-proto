package proto

// Map of GraphQL simplified names to proto names.
var GraphQLTypes = map[string]string{
	"UserPublicIdentity":    "github.com/fuserobotics/proto/auth.UserPublicIdentity",
	"CertChain":             "github.com/fuserobotics/proto/common.CertChain",
	"DeviceNetworkTemplate": "github.com/fuserobotics/proto/device.DeviceNetworkTemplate",
	"WifiNetwork":           "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork",
	"UserPrivateData":       "github.com/fuserobotics/proto/auth.UserPrivateData",
	"GeoLocation":           "github.com/fuserobotics/proto/common.GeoLocation",
	"IPRange":               "github.com/fuserobotics/proto/common.IPRange",
	"Device":                "github.com/fuserobotics/proto/device.Device",
	"DeviceIdentity":        "github.com/fuserobotics/proto/device.Device_DeviceIdentity",
	"DeviceStatus":          "github.com/fuserobotics/proto/device.DeviceStatus",
	"DeviceInterfaceConfig": "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig",
	"WifiConfig":            "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig",
	"DeviceNetworkSettings": "github.com/fuserobotics/proto/device.Device_DeviceNetworkSettings",
	"KVGDeviceSubKeys":      "github.com/fuserobotics/proto/device.KVGDeviceSubKeys",
	"ScryptParams":          "github.com/fuserobotics/proto/auth.UserPrivateData_ScryptParams",
	"IPAddress":             "github.com/fuserobotics/proto/common.IPAddress",
	"Region":                "github.com/fuserobotics/proto/region.Region",
	"GetUnixTimeResponse":   "github.com/fuserobotics/proto/timesync.GetUnixTimeResponse",
}
