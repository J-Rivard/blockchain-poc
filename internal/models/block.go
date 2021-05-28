package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Block struct {
	PreviousHash string
	Hash         string
	Txn          *Transaction
}

const requiredZeros = 6

func NewBlock(previousBlock *Block, txn *Transaction) (*Block, error) {
	block := &Block{}

	if previousBlock != nil {
		block = &Block{
			PreviousHash: previousBlock.Hash,
			Txn:          txn,
		}
	} else {
		block = &Block{
			Txn: txn,
		}
	}

	block.mineBlock()

	return block, nil
}

func GenerateGenesisBlock(publicKey, privateKey string) *Block {
	return &Block{}
}

func (b *Block) mineBlock() error {
	nonce := b.calcNonce()

	blockString, err := b.blockAsString(nonce)
	if err != nil {
		return err
	}

	b.Hash = Hash(blockString)

	return nil
}

func (b *Block) calcNonce() string {
	nonce := "I AM A TEAPOT SHORT AND STOUT"
	count := 0

	for !b.isValidNonce(nonce) {
		if count%100000 == 0 {
			fmt.Print(".")
		}

		nonce = IncrementString(nonce)
		count += 1
	}
	fmt.Println()

	return nonce
}

func (b *Block) isValidNonce(nonce string) bool {
	blockString, err := b.blockAsString(nonce)
	if err != nil {
		fmt.Println(err)
		return false
	}
	hashValue := Hash(blockString)

	goalString := strings.Repeat("0", requiredZeros)

	return strings.HasPrefix(hashValue, goalString)
}

func (b *Block) blockAsString(nonce string) (string, error) {

	txnBytes, err := json.Marshal(b.Txn)
	if err != nil {
		return "", err
	}

	return string(txnBytes) + b.PreviousHash + nonce, nil
}

func (b *Block) Print() {
	fmt.Printf("\nTxn_From: %sn\nTxn_To: %s\nTxn_Amt:%f\nPrevious_Hash:%s\nHash:0x%s\n\n", b.Txn.From, b.Txn.To, b.Txn.Amount, b.PreviousHash, b.Hash)
}
