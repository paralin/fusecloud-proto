package common

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var RootDomain string = "r.fusebot.io"

func RegionRootDomain(regionId string) string {
	return fmt.Sprintf("%s.%s", regionId, RootDomain)
}

func DeviceRootDomain(devHostname, regionId string) string {
	return fmt.Sprintf("%s.%s", devHostname, RegionRootDomain(regionId))
}

// Prepends a / for KVGossip root from FQDN.
func KVGossipFromFQDN(fqdn string) string {
	return fmt.Sprintf("/%s", fqdn)
}

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

func (loc *GeoLocation) Validate() error {
	if loc.Latitude < -90 || loc.Latitude > 90 {
		return errors.New("Latitude is invalid.")
	}
	if loc.Longitude > 180 || loc.Longitude < -180 {
		return errors.New("Longitude is invalid.")
	}
	return nil
}

// Really stupid implementation, will make something better later
func FromIPString(ip string) *IPAddress {
	parts := strings.Split(ip, ".")
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

func FromIPAddr(ip net.IP) *IPAddress {
	return FromIPString(ip.String())
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
