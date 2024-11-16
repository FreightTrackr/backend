package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

// ReadPrivateKeyFromFile reads an RSA private key from a file
func ReadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	var privateKey *rsa.PrivateKey
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(private.pem): %v", err)
	}
	privateBlock, _ := pem.Decode(privateKeyBytes)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		log.Fatalf("x509.ParsePKCS1PrivateKey: %v", err)
	}

	return privateKey, nil
}

func ReadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(public.pem): %v", err)
	}
	publicBlock, _ := pem.Decode(publicKeyBytes)
	pub, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		log.Fatalf("x509.ParsePKIXPublicKey: %v", err)
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("not ok: %v", err)
	}
	return publicKey, nil
}
