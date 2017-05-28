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
	"reflect"
	"strings"
	"time"

	"github.com/fuserobotics/proto/common"
)

type KVGDeviceIdentifier struct {
	Hostname string
	Region   string
}

type KVGDeviceSubKeys struct {
	DeviceInfo string
}

var DeviceKeyUsage x509.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment | x509.KeyUsageKeyEncipherment
var DeviceExtKeyUsage []x509.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageEmailProtection}
var KVGDeviceSubKeysBase = KVGDeviceSubKeys{
	DeviceInfo: "",
}

func ParseFqdn(fqdn string) (hostname string, region string, err error) {
	fqdnParts := strings.Split(fqdn, ".") // 2 + domainParts
	fqdnLen := len(fqdnParts)
	domainParts := strings.Split(common.RootDomain, ".") // 3
	domainLen := len(domainParts)
	if fqdnLen != 2+domainLen {
		return "", "", errors.New("Unexpected domain parts length.")
	}
	return fqdnParts[0], fqdnParts[1], nil
}

func ParseKVGKey(key string) *KVGDeviceIdentifier {
	parts := strings.Split(key, "/")
	// Expect leading / - /rv1.av1.r.fusebot.io
	if len(parts) < 2 {
		return nil
	}
	hostname := parts[1]
	if !strings.HasSuffix(hostname, common.RootDomain) {
		return nil
	}

	// rootDomainParts.len = 3
	rootDomainParts := strings.Split(common.RootDomain, ".")
	// hostnameParts.len = 5
	hostnameParts := strings.Split(hostname, ".")
	if (len(hostnameParts) - 2) != len(rootDomainParts) {
		return nil
	}
	return &KVGDeviceIdentifier{
		Hostname: hostnameParts[0],
		Region:   hostnameParts[1],
	}
}

func (d *Device) BuildKVGPaths() *KVGDeviceSubKeys {
	droot := common.KVGossipFromFQDN(d.DefaultFqdn())

	res := &KVGDeviceSubKeys{}
	vf := reflect.ValueOf(KVGDeviceSubKeysBase)

	rvf := reflect.ValueOf(res)
	rvfv := rvf.Elem()

	for i := 0; i < rvfv.NumField(); i++ {
		fields := vf.Field(i)
		fieldt := rvfv.Field(i)
		r := fmt.Sprintf("%s%s", droot, fields.Interface().(string))
		fieldt.SetString(r)
	}

	return res
}

func (d *Device) GenerateFqdn(domain string) string {
	return common.DeviceRootDomain(d.Hostname, d.Region)
}

func (d *Device) DefaultFqdn() string {
	return d.GenerateFqdn(common.RootDomain)
}

func (d *Device) Validate() error {
	hnlen := len(d.Hostname)
	if hnlen < 2 || hnlen > 20 {
		return errors.New("Hostname length is invalid.")
	}

	if d.Region == "" {
		return errors.New("Region must be specified.")
	}

	if d.NetworkSettings == nil {
		return errors.New("Network settings must be specified")
	}

	return nil
}

// Checks if a new value is OK (only runtime-changeable things changed).
func (old *Device) ValidateRuntimeMutation(newVal *Device) error {
	nfqdn := newVal.DefaultFqdn()
	ofqdn := old.DefaultFqdn()

	if err := newVal.Validate(); err != nil {
		return err
	}

	if nfqdn != ofqdn {
		return errors.New("Mutation would change fqdn (not allowed).")
	}

	if old.Identity.PublicKey != newVal.Identity.PublicKey {
		return errors.New("Public key would change (not allowed).")
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
/* Devices domain MUST be something like devices.fuserobotics.com */
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
		Organization:       []string{"FUSERobotics"},
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

func DeviceCertExpiryTime(from time.Time) time.Time {
	return from.Add(time.Duration(375) * 24 * time.Hour)
}
