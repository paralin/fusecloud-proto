package device

import (
	"github.com/fuserobotics/proto/common"
	"testing"
)

func TestParseFqdn(t *testing.T) {
	fqdn := "c2.ca1." + common.RootDomain
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

func TestBuildKVGPaths(t *testing.T) {
	dev := &Device{
		Hostname: "dc",
		Region:   "sim",
	}
	kp := dev.BuildKVGPaths()

	if kp.DeviceInfo != "/dc.sim.r.fusebot.io" {
		t.Fatalf(kp.DeviceInfo)
		t.Fail()
	}
}
