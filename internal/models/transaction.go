package models

import (
	"fmt"

	"github.com/J-Rivard/blockchain-poc/internal/pki"
)

type Transaction struct {
	From   string
	To     string
	Amount float64

	Signature string
}

func NewTxn(from, to, privateKey string, amount float64) *Transaction {
	txn := &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}

	txn.Signature = pki.Encrypt(txn.hash(), privateKey)

	return txn
}

func (t *Transaction) isValidSignature() bool {
	if t.isGenesisTransaction() {
		return true
	}

	return pki.IsValidSignature(t.hash(), t.Signature, t.From)
}

func (t *Transaction) isGenesisTransaction() bool {
	return t.From == ""
}

func (t *Transaction) hash() string {
	return fmt.Sprintf("%s%s%f", t.From, t.To, t.Amount)
}
