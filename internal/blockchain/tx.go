package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)


type Transaction struct{
	ID  []byte
	From string
	To string
	Amount int
	Type string
}

func(tx *Transaction) Hash() []byte{
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_=encoder.Encode(tx)
	hash := sha256.Sum256(buffer.Bytes())
	return hash[:]
}

func NewTransaction(from, to string, amount int, txType string) *Transaction{
	tx := &Transaction{
		From: from,
		To: to,
		Amount: amount,
		Type: txType,
	}
	tx.ID = tx.Hash()
	return tx
}