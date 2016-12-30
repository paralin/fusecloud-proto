package auth

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"testing"
)

var TestUsername string = "batman"
var TestPassword string = "NaNNaNNaNNaN"
var TestKey *rsa.PrivateKey

func init() {
	nk, _ := rsa.GenerateKey(rand.Reader, 2048)
	TestKey = nk
}

func TestScryptAuth(t *testing.T) {
	pd := &UserPrivateData{}
	if err := pd.ApplyPassword(TestUsername, TestPassword, TestKey); err != nil {
		t.Fatal(err.Error())
	}
	if err := pd.Validate(); err != nil {
		t.Fatal(err.Error())
	}
	npkey, err := pd.UsePassword(TestUsername, TestPassword)
	if err != nil {
		t.Fatal(err.Error())
	}
	if bytes.Compare(x509.MarshalPKCS1PrivateKey(TestKey), x509.MarshalPKCS1PrivateKey(npkey)) != 0 {
		t.Fatal("Keys did not match.")
	}
}
