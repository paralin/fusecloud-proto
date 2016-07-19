package device

var DevicesTable string = "devices"
var RegionTableTemplate map[string]interface{} = map[string]interface{}{
	"hostname": "",
	"region":   "",
	"proto":    []byte{},
}
