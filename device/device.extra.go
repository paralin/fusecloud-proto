package device

type DeviceWithStatus struct {
	Device *Device      `json:"device"`
	Status DeviceStatus `json:"status"`
}
