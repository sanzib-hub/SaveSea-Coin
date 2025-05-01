package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
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
	Difficulty int
}




func NewBlock(transactions []*Transaction, prevHash []byte, index, difficulty int) *Block{
	block := &Block{
		Index: index,
		Timestamp: time.Now().Unix(),
		Transactions: transactions,
		PrevHash: prevHash,
		Nonce: 0,
		Difficulty: difficulty,
	}
	block.MerkleRoot = block.buildMerkleRoot()
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() []byte{
	data:= bytes.Join(
		[][]byte{
			IntToHex(int64(b.Index)),
			IntToHex(b.Timestamp),
			b.PrevHash,
			b.MerkleRoot,
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

	combined:=bytes.Join(txHashes, []byte{})
	hash:= sha256.Sum256(combined)
	return hash[:]

}
func(b *Block) buildMerkleRoot() []byte{
	if len(b.Transactions)==0{
		return nil
	}
	var merkleTree [][]byte
	for _,tx := range b.Transactions{
		merkleTree = append(merkleTree, tx.Hash())
	}
	for len(merkleTree)>1{
		if len(merkleTree)%2 !=0{
			merkleTree = append(merkleTree, merkleTree[len(merkleTree)-1])

		}
		var newLevel [][]byte

		for i:=0;i<len(merkleTree); i +=2{
			combined:= append(merkleTree[i],merkleTree[i+1]...)
			hash:=sha256.Sum256(combined)
			newLevel = append(newLevel, hash[:])
		}
		merkleTree = newLevel
	}
	if len(merkleTree)==1{
		return merkleTree[0]
	}
	return nil
}

func(b *Block) Mine(){
	target := bytes.Repeat([]byte{0}, b.Difficulty)
	for {
		b.Hash = b.calculateHash()
		if bytes.HasPrefix(b.Hash, target){

			return
		}
		b.Nonce++
	}
}

func (b *Block) Serialize() ([]byte, error){
	// use bytes.Buffer to memory manage
	buf:= new(bytes.Buffer)
	// var buf bytes.Buffer
	encoder := gob.NewEncoder(buf)
	if err :=encoder.Encode(b); err !=nil{
		return nil, fmt.Errorf("failed to serialize block: %w", err)
	}

	return buf.Bytes(),nil

}

func DeserializeBlock(data []byte) (*Block, error){
	if len(data) == 0 {
		return nil, errors.New("cannot deserialize empty data")
	}
	var block Block
	reader := bytes.NewReader(data)
	
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(&block); err != nil {
		return nil, fmt.Errorf("failed to deserialize block: %w", err)
	}
	
	return &block, nil
}