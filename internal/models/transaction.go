package models

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/J-Rivard/blockchain-poc/internal/pki"
)

type Transaction struct {
	From   string
	To     string
	Amount float64

	Signature string
}

func NewTxn(from, to string, privateKey *rsa.PrivateKey, amount float64) (*Transaction, error) {
	t := &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	var err error

	t.Signature, err = pki.Sign(privateKey, t.hash())
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Transaction) isValidSignature(signature string) bool {
	if t.isGenesisTransaction() {
		return true
	}

	publicBlock, _ := pem.Decode([]byte(t.From))
	publicKey, err := x509.ParsePKCS1PublicKey(publicBlock.Bytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return pki.Verify(publicKey, []byte(signature), t.hash())
}

func (t *Transaction) isGenesisTransaction() bool {
	return t.From == ""
}

func (t *Transaction) hash() string {
	return fmt.Sprintf("%s%s%f", t.From, t.To, t.Amount)
}
