package device

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
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
	pkd, err := x509.MarshalPKIXPublicKey(pkey.Public())
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
		CommonName:         fmt.Sprintf("%s.%s.%s", ddev.Hostname, ddev.Region, devicesDomain),
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

func DeviceCertExpiryTime(from time.Time) time.Time {
	return from.Add(time.Duration(375) * 24 * time.Hour)
}
