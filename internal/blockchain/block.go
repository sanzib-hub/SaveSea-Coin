package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)
type Block struct{
	Index int
	Timestamp int64
	Transactions []*Transaction
	PrevHash []byte
	Hash []byte
	Nonce int
	MerkleRoot  []byte
}




func NewBlock(transactions []*Transaction, prevHash []byte, index int) *Block{
	block := &Block{
		Index: index,
		Timestamp: time.Now().Unix(),
		Transactions: transactions,
		PrevHash: prevHash,
		Nonce: 0,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() []byte{
	data:= bytes.Join(
		[][]byte{
			IntToHex(int64(b.Index)),
			IntToHex(b.Timestamp),
			b.PrevHash,
			b.hashTransactions(),
			IntToHex(int64(b.Nonce)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	return hash[:]
}
func(b *Block) hashTransactions() []byte{
	var txHashes [][]byte

	for _, tx:= range b.Transactions{
		txHashes = append(txHashes, tx.Hash())
	}
}
