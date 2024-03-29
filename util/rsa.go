package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var PemDecodeErr error = errors.New("Pem decoding failed.")
var UnknownPublicKeyErr error = errors.New("Key is not an RSA public key.")

func ParsePublicKey(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, PemDecodeErr
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, UnknownPublicKeyErr
	}
}

func ParsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, PemDecodeErr
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func MarshalPrivateKey(pkey *rsa.PrivateKey) []byte {
	blk := pem.Block{}
	blk.Type = "PRIVATE KEY"
	blk.Bytes = x509.MarshalPKCS1PrivateKey(pkey)

	return pem.EncodeToMemory(&blk)
}

func ComparePublicKey(k1 *rsa.PublicKey, k2 *rsa.PublicKey) bool {
	return k1.E == k2.E && k1.N.Cmp(k2.N) == 0
}

func ComparePublicKeyIB(k1 *rsa.PublicKey, k2b []byte) bool {
	k2, err := ParsePublicKey(k2b)
	if err != nil {
		return false
	}
	return ComparePublicKey(k1, k2)
}

func MarshalPublicKey(k *rsa.PublicKey) ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(k)
	if err != nil {
		return nil, err
	}
	blk := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: data,
	}
	return pem.EncodeToMemory(&blk), nil
}
