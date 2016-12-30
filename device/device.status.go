package device

type DeviceStatus struct {
	*Device
	// Serf status
	SerfStatus string
}
