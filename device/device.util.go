package device

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"

	"encoding/pem"
	"errors"
	"fmt"
	"net"

	"github.com/synrobo/proto/common"
)

var DevicesDomain string = "devices.synrobo.com"

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

func (di *Device_DeviceIdentity) GenerateKey() error {
	pkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	di.Pkey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pkey),
	}))
	return nil
}

func (ddev *Device_DeviceIdentity) EnsureKey() error {
	if ddev.Pkey == "" {
		return ddev.GenerateKey()
	}
	return nil
}

func (ddev *Device_DeviceIdentity) ParsePkey() (*rsa.PrivateKey, error) {
	if ddev.Pkey == "" {
		return nil, errors.New("Device has no private key.")
	}
	blk, _ := pem.Decode([]byte(ddev.Pkey))
	if blk == nil {
		return nil, errors.New("Private key pem failed to parse")
	}
	if blk.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("Pkey pem type %s is not RSA PRIVATE KEY.", blk.Type)
	}
	return x509.ParsePKCS1PrivateKey(blk.Bytes)
}

func (dev *Device) EnsureIdentity() error {
	if dev.Identity == nil {
		dev.Identity = &Device_DeviceIdentity{}
	}
	return dev.Identity.EnsureKey()
}

/* Generates a CSR and adds it to the beginning of the list. */
/* Devices domain MUST be something like devices.synrobo.com */
func (ddev *Device) AddCert(devicesDomain, regionName string) error {
	if err := ddev.EnsureIdentity(); err != nil {
		return err
	}

	pkey, err := ddev.Identity.ParsePkey()
	if err != nil {
		return err
	}

	newCert := &common.CertChain{
		CsrWaiting: true,
	}

	// generate csr
	reqTmpl := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            []string{"United States"},
			Province:           []string{"California"},
			Locality:           []string{regionName},
			Organization:       []string{"Synergy Robotics"},
			OrganizationalUnit: []string{"devices:" + ddev.Region},
			CommonName:         fmt.Sprintf("%s.%s.%s", ddev.Hostname, ddev.Region, devicesDomain),
		},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, reqTmpl, pkey)
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

func (dev *Device) EnsureIdentityAndCert(regionName string) error {
	if err := dev.EnsureIdentity(); err != nil {
		return err
	}
	if len(dev.Identity.Chain) == 0 {
		return dev.AddCert(DevicesDomain, regionName)
	}
	return nil
}
