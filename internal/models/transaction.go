package models

import "fmt"

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

	return txn
}

func (t *Transaction) isGenesisTransaction() bool {
	return t.From == ""
}

func (t *Transaction) hash() string {
	return fmt.Sprintf("%s%s%f", t.From, t.To, t.Amount)
}
