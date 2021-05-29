package pki

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func Sign(privateKey *rsa.PrivateKey, plaintext string) (string, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(plaintext))
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)

	signed, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

func Verify(publicKey *rsa.PublicKey, signature []byte, plaintext string) bool {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(plaintext))
	if err != nil {
		return false
	}
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return false
	}

	return true
}
