package device

import (
	"testing"
)

func TestParseFqdn(t *testing.T) {
	fqdn := "c2.ca1." + DevicesDomain
	hostname, region, err := ParseFqdn(fqdn)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if hostname != "c2" {
		t.Errorf("Hostname is wrong, %s != c2", hostname)
		t.Fail()
		return
	}

	if region != "ca1" {
		t.Errorf("Region is wrong, %s != ca1", region)
		t.Fail()
		return
	}
}
