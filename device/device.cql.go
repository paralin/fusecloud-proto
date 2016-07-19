package device

var DevicesTable string = "devices"
var DevicesTableTemplate map[string]interface{} = map[string]interface{}{
	"hostname": "",
	"region":   "",
	"proto":    []byte{},
}
