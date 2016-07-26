package common

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (ip *IPAddress) Validate() error {
	if ip == nil || len(ip.Address) != 4 {
		return errors.New("IP address invalid.")
	}
	return nil
}

func (ipr *IPRange) Validate() error {
	if err := ipr.Ip.Validate(); err != nil {
		return err
	}
	if ipr.Plen > 32 {
		return fmt.Errorf("Prefix length %d is invalid.", ipr.Plen)
	}
	return nil
}

func (ip *IPAddress) IPString() (string, error) {
	if err := ip.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d", ip.Address[0], ip.Address[1], ip.Address[2], ip.Address[3]), nil
}

func (ip *IPAddress) ToIPAddr() (net.IP, error) {
	ipstr, err := ip.IPString()
	if err != nil {
		return nil, err
	}
	return net.ParseIP(ipstr), nil
}

// Really stupid implementation, will make something better later
func FromIPAddr(ip net.IP) *IPAddress {
	parts := strings.Split(ip.String(), ".")
	res := &IPAddress{
		Address: []uint32{},
	}
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		res.Address = append(res.Address, uint32(i))
	}
	return res
}

func (ipr *IPRange) CIDRString() (string, error) {
	ips, err := ipr.Ip.IPString()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%d", ips, ipr.Plen), nil
}

func (ipr *IPRange) BuildCIDRInfo() (net.IP, *net.IPNet, error) {
	cidrStr, err := ipr.CIDRString()
	if err != nil {
		return nil, nil, err
	}
	return net.ParseCIDR(cidrStr)
}

func (ipr *IPRange) ValidateDeviceIP(ip net.IP) error {
	_, cidr, err := ipr.BuildCIDRInfo()
	if err != nil {
		return err
	}
	if !cidr.Contains(ip) {
		return errors.New("Cidr does not contain IP.")
	}
	return nil
}
