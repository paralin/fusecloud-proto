package device

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"strings"
	"time"

	"encoding/pem"
	"errors"
	"fmt"
	"net"

	"github.com/synrobo/proto/common"
)

var DevicesDomain string = "devices.synrobo.com"
var DeviceKeyUsage x509.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment
var DeviceExtKeyUsage []x509.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageEmailProtection}

func ParseFqdn(fqdn string) (hostname string, region string, err error) {
	fqdnParts := strings.Split(fqdn, ".") // 2 + domainParts
	fqdnLen := len(fqdnParts)
	domainParts := strings.Split(DevicesDomain, ".") // 3
	domainLen := len(domainParts)
	if fqdnLen != 2+domainLen {
		return "", "", errors.New("Unexpected domain parts length.")
	}
	return fqdnParts[0], fqdnParts[1], nil
}

func (d *Device) GenerateFqdn(domain string) string {
	return fmt.Sprintf("%s.%s.%s", d.Hostname, d.Region, domain)
}

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
	if err := dn.Ip.Validate(); err != nil {
		return nil, err
	}
	return dn.Ip.ToIPAddr()
}

func (di *Device_DeviceIdentity) GenerateKey() (*rsa.PrivateKey, error) {
	pkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	pkd, err := x509.MarshalPKIXPublicKey(&pkey.PublicKey)
	if err != nil {
		return nil, err
	}
	di.PublicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pkd,
	}))
	return pkey, nil
}

func (ddev *Device_DeviceIdentity) ParsePublicKey() (*rsa.PublicKey, error) {
	if ddev.PublicKey == "" {
		return nil, errors.New("Device has no private key.")
	}
	blk, _ := pem.Decode([]byte(ddev.PublicKey))
	if blk == nil {
		return nil, errors.New("Private key pem failed to parse")
	}
	if blk.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("Public key pem type %s is not PUBLIC KEY.", blk.Type)
	}
	pkr, err := x509.ParsePKIXPublicKey(blk.Bytes)
	if err != nil {
		return nil, err
	}
	return pkr.(*rsa.PublicKey), nil
}

/* Generates a CSR and adds it to the beginning of the list. */
/* Devices domain MUST be something like devices.synrobo.com */
func (ddev *Device) AddCert(devicesDomain, regionName string, pkey *rsa.PrivateKey) error {
	if ddev.Identity == nil {
		ddev.Identity = &Device_DeviceIdentity{}
	}

	newCert := &common.CertChain{
		CsrWaiting: true,
	}

	csr, err := ddev.BuildCertificateRequest(regionName, devicesDomain, pkey)
	if err != nil {
		return err
	}

	pemd := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	})

	newCert.Csr = string(pemd)
	ddev.Identity.Chain = append([]*common.CertChain{newCert}, ddev.Identity.Chain...)
	return nil
}

func (ddev *Device) BuildSubject(regionName string, devicesDomain string) pkix.Name {
	return pkix.Name{
		Country:            []string{"United States"},
		Province:           []string{"California"},
		Locality:           []string{regionName},
		Organization:       []string{"Synergy Robotics"},
		OrganizationalUnit: []string{"devices:" + ddev.Region},
		CommonName:         ddev.GenerateFqdn(devicesDomain),
	}
}

func (ddev *Device) BuildCertificateRequest(regionName string, devicesDomain string, pkey *rsa.PrivateKey) ([]byte, error) {
	// generate csr
	reqTmpl := &x509.CertificateRequest{
		Subject:            ddev.BuildSubject(regionName, devicesDomain),
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	return x509.CreateCertificateRequest(rand.Reader, reqTmpl, pkey)
}

func (ns *Device_DeviceInterfaceConfig) BuildSystemdNetworkdFile() string {
	var res bytes.Buffer
	res.WriteString("[Match]\nName=")
	res.WriteString(ns.Devname)
	res.WriteString("\n\n[Network]\n")

	var ips string
	var ipsg string

	if ns.Ip != nil {
		ips, _ = ns.Ip.IPString()
	} else {
		ips = ""
	}
	if ips != "" {
		res.WriteString("Address=")
		res.WriteString(ips)
		res.WriteString("/24\n")
	}

	if ns.GatewayIp != nil {
		ipsg, _ = ns.GatewayIp.IPString()
	}
	if ipsg != "" {
		res.WriteString("Gateway=")
		res.WriteString(ipsg)
		res.WriteString("\n")
	}

	if ips == "" && ipsg == "" {
		res.WriteString("DHCP=ipv4\n")
	}

	if len(ns.Dns) > 0 {
		for _, item := range ns.Dns {
			if item == nil {
				continue
			}

			dnsstr, err := item.IPString()
			if err != nil {
				continue
			}
			res.WriteString("DNS=")
			res.WriteString(dnsstr)
			res.WriteString("\n")
		}
	}

	return res.String()
}

func (wif *Device_DeviceInterfaceConfig_WifiConfig) HasAnyAdhoc() bool {
	for _, net := range wif.Network {
		if net.NetType == Device_DeviceInterfaceConfig_WifiConfig_ADHOC {
			return true
		}
	}
	return false
}

func (wif *Device_DeviceInterfaceConfig_WifiConfig) BuildWpaSupplicantFile() string {
	var res bytes.Buffer

	res.WriteString("ctrl_interface=/var/run/wpa_supplicant\n")
	if wif.HasAnyAdhoc() {
		res.WriteString("# There is at least 1 adhoc network, use ap_scan and eapol:\n")
		res.WriteString("eapol_version=1\n")
		res.WriteString("ap_scan=2\n")
	} else {
		res.WriteString("ap_scan=1\n")
		res.WriteString("fast_reauth=1\n")
	}

	if wif.ExtraOptions != "" {
		res.WriteString("\n# User-supplied extra options\n")
		res.WriteString(wif.ExtraOptions)
		res.WriteString("\n")
	}

	res.WriteString("\n# Networks generated by devman\n")
	for _, net := range wif.Network {
		res.WriteString("network={\n")
		res.WriteString(fmt.Sprintf("  ssid=\"%s\"\n", net.Ssid))
		switch net.NetType {
		case Device_DeviceInterfaceConfig_WifiConfig_ADHOC:
			res.WriteString("  mode=1\n")
			if net.Psk != "" {
				res.WriteString("  proto=WPA\n")
				res.WriteString("  key_mgmt=WPA-NONE\n")
				res.WriteString("  pairwise=NONE\n")
				res.WriteString("  group=TKIP\n")
			} else {
				res.WriteString("  proto=RSN\n")
				res.WriteString("  key_mgmt=NONE\n")
			}
			if net.Frequency == "" {
				// Adhoc networks need a frequency, default to 2432
				net.Frequency = "2432"
			}
		case Device_DeviceInterfaceConfig_WifiConfig_AP:
			if net.Proto != "" {
				res.WriteString(fmt.Sprintf("  proto=%s\n", net.Proto))
			}
			if net.KeyMgmt != "" {
				res.WriteString(fmt.Sprintf("  key_mgmt=%s\n", net.KeyMgmt))
			}
			if net.Pairwise != "" {
				res.WriteString(fmt.Sprintf("  pairwise=%s\n", net.Pairwise))
			}
			if net.Group != "" {
				res.WriteString(fmt.Sprintf("  group=%s\n", net.Group))
			}
		}
		if net.Frequency != "" {
			res.WriteString(fmt.Sprintf("  frequency=%s\n", net.Frequency))
		}
		if net.Psk != "" {
			res.WriteString("  psk=")
			if !net.PskEncoded {
				res.WriteString("\"")
			}
			res.WriteString(net.Psk)
			if !net.PskEncoded {
				res.WriteString("\"")
			}
			res.WriteString("\n")
		}
		if net.ExtraOptions != "" {
			res.WriteString("  # User supplied extra options\n")
			res.WriteString(fmt.Sprintf("  %s\n", strings.Replace(net.ExtraOptions, "\n", "\n  ", -1)))
		}
		res.WriteString("}\n")
	}
	return res.String()
}

func DeviceCertExpiryTime(from time.Time) time.Time {
	return from.Add(time.Duration(375) * 24 * time.Hour)
}
