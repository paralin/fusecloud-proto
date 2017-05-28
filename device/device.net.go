package device

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc64"
	"math/rand"
	"strings"

	"github.com/satori/go.uuid"
)

// DeviceConnectionTypeIds pairs DeviceConnectionType with NetworkManager ID string
var DeviceConnectionTypeIds = map[DeviceConnectionType]string{
	DeviceConnectionType_CONN_ETH:  "ethernet",
	DeviceConnectionType_CONN_WIFI: "wifi",
}

// DeviceWifiModeIds pairs DeviceConnection_WifiConfig_WifiMode with NetworkManager ID string
var DeviceWifiModeIds = map[DeviceConnection_WifiConfig_WifiMode]string{
	DeviceConnection_WifiConfig_WIFI_ADHOC: "adhoc",
	DeviceConnection_WifiConfig_WIFI_INFRA: "infrastructure",
}

// UUID builds a UUID based on the JSON representation of the connection.
func (c *DeviceConnection) ToUUID() string {
	dat, _ := json.Marshal(c)
	r := rand.New(rand.NewSource(int64(crc64.Checksum(dat, crc64.MakeTable(crc64.ISO)))))
	u := uuid.UUID{}
	r.Read(u[:])
	u.SetVersion(4)
	u.SetVariant()
	return u.String()
}

// BuildKeyfile builds a NetworkManager keyfile from the connection.
func (c *DeviceConnection) BuildKeyfile() string {
	var output bytes.Buffer
	w := func(s string, args ...interface{}) { output.WriteString(fmt.Sprintf(s, args...)) }

	uuid := c.ToUUID()
	w("[connection]\n")
	w("id=%s\n", uuid)
	w("uuid=%s\n", uuid)
	w("type=%s\n", DeviceConnectionTypeIds[c.Type])
	w("autoconnect=%v\n", c.AutoConnect)
	if c.AutoConnect {
		if c.AutoConnectPriority > 0 {
			w("autoconnect-priority=%d\n", c.AutoConnectPriority)
		}
		if c.AutoConnectRetries > 0 {
			w("autoconnect-retries=%d\n", c.AutoConnectRetries)
		}
	}
	if c.Interface != "" {
		w("interface-name=%s\n", c.Interface)
	}

	if c.Type == DeviceConnectionType_CONN_WIFI && c.Wifi != nil {
		wifi := c.Wifi
		w("\n[wifi]\n")
		w("mode=%s\n", DeviceWifiModeIds[wifi.Mode])
		w("ssid=%s\n", wifi.Ssid)

		if wifi.Security != nil && wifi.Security.Mode != DeviceConnection_WifiSecurity_WIFI_SEC_NONE {
			w("\n[wifi-security]\n")

			switch wifi.Security.Mode {
			case DeviceConnection_WifiSecurity_WIFI_SEC_PERSONAL:
				w("key-mgmt=wpa-psk\n")
				w("psk=%s\n", wifi.Security.Password)
			case DeviceConnection_WifiSecurity_WIFI_SEC_ENTERPRISE:
				w("key-mgmt=wpa-eap\n")
				w("\n[802-1x]\n")
				w("eap=peap;\n")
				w("identity=%s\n", wifi.Security.Username)
				w("password=%s\n", wifi.Security.Password)
				w("phase2-auth=mschapv2\n")
			}
		}
	} else if c.Type == DeviceConnectionType_CONN_ETH {
		w("\n[ethernet]\n")
		w("mac-address-blacklist=")
	}

	w("\n[ipv6]\n")
	w("addr-gen-mode=stable-privacy\n")
	w("method=auto\n")

	if c.Ipv4 == nil {
		c.Ipv4 = &DeviceConnection_IPV4Config{}
	}

	w("\n[ipv4]\n")
	w("method=")
	if len(c.Ipv4.Address) > 0 {
		w("auto\n")
	} else {
		w("manual\n")
		addri := 1
		for _, addr := range c.Ipv4.Address {
			if addr.Range == nil {
				continue
			}
			cs, er := addr.Range.CIDRString()
			if er != nil {
				continue
			}
			w("address%d=%s", addri, cs)
			if addr.Gateway != nil {
				gs, err := addr.Gateway.IPString()
				if err == nil {
					w(",%s", gs)
				}
			}
			w("\n")
			addri++
		}
	}
	w("dns=")
	for _, search := range c.Ipv4.Dns {
		str, er := search.IPString()
		if er != nil {
			continue
		}
		w("%s;", str)
	}
	w("\ndns-search=")
	for _, search := range c.Ipv4.DnsSearch {
		w(strings.TrimSpace(search))
		w(";")
	}
	w("\n")
	if c.Ipv4.DisableAutoDns {
		w("ignore-auto-dns=true\n")
	}

	return output.String()
}
