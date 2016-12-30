package region

import (
	"errors"
	"github.com/fuserobotics/proto/common"
	"strings"
)

func (r *Region) Validate() error {
	if r.Id == "" {
		return errors.New("ID must not be null.")
	}
	if r.Name == "" {
		return errors.New("Name must not be null.")
	}
	if r.IpRange == nil {
		return errors.New("IP range must not be null.")
	}
	if err := r.IpRange.Validate(); err != nil {
		return err
	}
	if r.Location == nil {
		return errors.New("Location must not be null.")
	}
	if err := r.Location.Validate(); err != nil {
		return err
	}
	return nil
}

func ParseFqdn(fqdn string) (region string, err error) {
	fqdnParts := strings.Split(fqdn, ".") // 1 + domainParts
	fqdnLen := len(fqdnParts)
	domainParts := strings.Split(common.RootDomain, ".") // 2
	domainLen := len(domainParts)
	if fqdnLen != 1+domainLen {
		return "", errors.New("Unexpected domain parts length.")
	}
	return fqdnParts[0], nil
}
