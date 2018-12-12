package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/andskur/experemental/services/blockchain/merkle_tree"
	"github.com/andskur/experemental/services/blockchain/txs"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     	int64
	Transactions	[]*txs.Transaction
	PrevBlockHash 	[]byte
	Hash          	[]byte
	Nonce         	int
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := merkle_tree.NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// NewBlock creates and returns Block
func NewBlock(transactions []*txs.Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *txs.Transaction) *Block {
	return NewBlock([]*txs.Transaction{coinbase}, []byte{})
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
