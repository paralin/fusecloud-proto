package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
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

func (pu *UserPrivateData_ScryptParams) Validate() error {
	if pu.N <= 1 || (pu.N&(pu.N-1)) != 0 {
		return errors.New("N must be a power of 2 greater than 1.")
	}
	if pu.P != 1 {
		return errors.New("P must be 1.")
	}
	if pu.R > 30 || pu.R < 1 {
		return errors.New("R must be 1 <= R <= 30")
	}
	return nil
}

func (pu *UserPrivateData) FillDefaults() {
	if pu.ScryptParams == nil {
		pu.ScryptParams = &UserPrivateData_ScryptParams{
			N: 16384,
			R: 8,
			P: 1,
		}
	}
}

func (pu *UserPrivateData) Validate() error {
	if pu.Strategy != UserPrivateData_PKEY_AES256 {
		return errors.New("Only AES256 is supported.")
	}
	// I computed 1191 length, but let's give some leeway.
	if len(pu.PrivateKey) > 1300 {
		return errors.New("Expected private key length to be less than 1300.")
	}
	if len(pu.PrivateKey) < 300 {
		return errors.New("Expected private key length to be at least 300.")
	}
	if pu.ScryptParams == nil {
		return errors.New("Scrypt params cannot be nil.")
	}
	if err := pu.ScryptParams.Validate(); err != nil {
		return err
	}

	return nil
}
