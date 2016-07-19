package device

import (
	"errors"
)

func (d *Device) Validate() error {
	hnlen := len(d.Hostname)
	if hnlen < 2 || hnlen > 20 {
		return errors.New("Hostname length is invalid.")
	}

	if d.Region == "" {
		return errors.New("Region must be specified.")
	}

	if d.ConsulSettings == nil {
		return errors.New("Consul settings must be specified.")
	}

	if d.NetworkSettings == nil {
		return errors.New("Network settings must be specified")
	}

	return nil
}
