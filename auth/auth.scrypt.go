package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func (pu *UserPrivateData_ScryptParams) Scrypt(password, salt []byte, rlen int) ([]byte, error) {
	return scrypt.Key(password, salt, int(pu.N), int(pu.R), int(pu.P), rlen)
}

func (pu *UserPrivateData) ApplyPassword(username, password string, pkey *rsa.PrivateKey) error {
	if username == "" || password == "" || pkey == nil {
		return errors.New("Username, password, and private key must be specified.")
	}

	// Fill scrypt params if necessary
	pu.FillDefaults()

	pkeym := x509.MarshalPKCS1PrivateKey(pkey)

	akey, iv, err := pu.ComputeScryptKeys(username, password)
	if err != nil {
		return err
	}

	cblock, err := aes.NewCipher(akey)
	if err != nil {
		return err
	}

	if cblock.BlockSize() != len(iv) {
		return fmt.Errorf("Block size %d != iv length %d", cblock.BlockSize(), len(iv))
	}

	encrypted := make([]byte, len(pkeym))
	cfbstream := cipher.NewCFBEncrypter(cblock, iv)
	cfbstream.XORKeyStream(encrypted, pkeym)

	pu.PrivateKey = encrypted
	return nil
}

func (pu *UserPrivateData) ComputeScryptKeys(username, password string) (akey []byte, iv []byte, err error) {
	akey, err = pu.ScryptParams.Scrypt([]byte(password), []byte(username), 32)
	if err != nil {
		return
	}

	iv, err = pu.ScryptParams.Scrypt([]byte(password), []byte(username), 16)
	if err != nil {
		return
	}

	return
}

func (pu *UserPrivateData) UsePassword(username, password string) (*rsa.PrivateKey, error) {
	if err := pu.Validate(); err != nil {
		return nil, err
	}

	akey, iv, err := pu.ComputeScryptKeys(username, password)
	if err != nil {
		return nil, err
	}

	cblock, err := aes.NewCipher(akey)
	if err != nil {
		return nil, err
	}

	if cblock.BlockSize() != len(iv) {
		return nil, fmt.Errorf("Block size %d != iv length %d", cblock.BlockSize(), len(iv))
	}

	pkeybytes := make([]byte, len(pu.PrivateKey))
	cfbstream := cipher.NewCFBDecrypter(cblock, iv)
	cfbstream.XORKeyStream(pkeybytes, pu.PrivateKey)
	return x509.ParsePKCS1PrivateKey(pkeybytes)
}
