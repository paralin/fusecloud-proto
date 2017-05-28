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

func GenerateFqdn(id string) string {
	return common.RegionRootDomain(id)
}

type KVGRegionIdentifier struct {
	Id string
}

func ParseKVGKey(key string) *KVGRegionIdentifier {
	parts := strings.Split(key, "/")
	// Expect leading / - /av1.r.fusebot.io
	if len(parts) < 2 {
		return nil
	}
	hostname := parts[1]
	if !strings.HasSuffix(hostname, common.RootDomain) {
		return nil
	}

	// rootDomainParts.len = 3
	rootDomainParts := strings.Split(common.RootDomain, ".")
	// hostnameParts.len = 4
	hostnameParts := strings.Split(hostname, ".")
	if (len(hostnameParts) - 1) != len(rootDomainParts) {
		return nil
	}
	return &KVGRegionIdentifier{
		Id: hostnameParts[0],
	}
}
