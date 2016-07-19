package common

import "fmt"

func (ipr *IPRange) Validate() error {
	if len(ipr.Ip) != 4 {
		return fmt.Errorf("Ip length %d != 4", len(ipr.Ip))
	}
	if ipr.Plen > 32 {
		return fmt.Errorf("Prefix length %d is invalid.", ipr.Plen)
	}
	return nil
}

func (ipr *IPRange) IPString() (string, error) {
	if err := ipr.Validate(); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d", ipr.Ip[0], ipr.Ip[1], ipr.Ip[2], ipr.Ip[3]), nil
}

func (ipr *IPRange) CIDRString() (string, error) {
	ips, err := ipr.IPString()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%d", ips, ipr.Plen), nil
}
