package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

const AuthDomain string = "u.fusebot.io"

func (user *UserPublicIdentity) GenerateFqdn(authDomain string) string {
	return fmt.Sprintf("%s.%s", user.Username, authDomain)
}

func ParseFqdn(fqdn string) (username string, err error) {
	fqdnParts := strings.Split(fqdn, ".") // 1 + domainParts
	fqdnLen := len(fqdnParts)
	domainParts := strings.Split(AuthDomain, ".") // 3
	domainLen := len(domainParts)
	if fqdnLen != 1+domainLen {
		return "", errors.New("Unexpected domain parts length.")
	}
	return fqdnParts[0], nil
}

func (usr *UserPublicIdentity) ParsePublicKey() (*rsa.PublicKey, error) {
	if usr.PublicKey == "" {
		return nil, errors.New("User has no public key.")
	}
	blk, _ := pem.Decode([]byte(usr.PublicKey))
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

func (user *UserPublicIdentity) BuildSubject(authDomain string) pkix.Name {
	return pkix.Name{
		Country:            []string{"United States"},
		Province:           []string{"California"},
		Locality:           []string{""},
		Organization:       []string{"FUSE Robotics"},
		OrganizationalUnit: []string{"users:" + user.Username},
		CommonName:         user.GenerateFqdn(authDomain),
	}
}

func UserCertExpiryTime(from time.Time) time.Time {
	return from.Add(time.Duration(375) * 24 * time.Hour)
}
