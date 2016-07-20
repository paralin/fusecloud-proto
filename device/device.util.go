package device

import (
	"errors"
	"fmt"
	"net"
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

func (dn *Device_DeviceNetworkSettings) ToIP() (net.IP, error) {
	if len(dn.Ip) < 4 {
		return nil, errors.New("IP length is less than 4.")
	}
	return net.ParseIP(fmt.Sprintf("%d.%d.%d.%d", dn.Ip[0], dn.Ip[1], dn.Ip[2], dn.Ip[3])), nil
}
