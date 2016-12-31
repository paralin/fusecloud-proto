package proto

// Map of GraphQL simplified names to proto names.
var GraphQLTypes = map[string]string{
	"CertChain":             "github.com/fuserobotics/proto/common.CertChain",
	"Device":                "github.com/fuserobotics/proto/device.Device",
	"DeviceIdentity":        "github.com/fuserobotics/proto/device.Device_DeviceIdentity",
	"DeviceInterfaceConfig": "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig",
	"DeviceNetworkSettings": "github.com/fuserobotics/proto/device.Device_DeviceNetworkSettings",
	"DeviceNetworkTemplate": "github.com/fuserobotics/proto/device.DeviceNetworkTemplate",
	"DeviceStatus":          "github.com/fuserobotics/proto/device.DeviceStatus",
	"GeoLocation":           "github.com/fuserobotics/proto/common.GeoLocation",
	"GetUnixTimeResponse":   "github.com/fuserobotics/proto/timesync.GetUnixTimeResponse",
	"IPAddress":             "github.com/fuserobotics/proto/common.IPAddress",
	"IPRange":               "github.com/fuserobotics/proto/common.IPRange",
	"KVGDeviceSubKeys":      "github.com/fuserobotics/proto/device.KVGDeviceSubKeys",
	"Region":                "github.com/fuserobotics/proto/region.Region",
	"ScryptParams":          "github.com/fuserobotics/proto/auth.UserPrivateData_ScryptParams",
	"UserPrivateData":       "github.com/fuserobotics/proto/auth.UserPrivateData",
	"UserPublicIdentity":    "github.com/fuserobotics/proto/auth.UserPublicIdentity",
	"WifiConfig":            "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig",
	"WifiNetwork":           "github.com/fuserobotics/proto/device.Device_DeviceInterfaceConfig_WifiConfig_WifiNetwork",
}
