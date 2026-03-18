package decentralizestorage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	merkletree "home/pacman/merkletree"
	"io"
	"log"
	"strings"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

// a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := merkletree.NewMerkleTree(transactions)
	return mTree.RootNode.Data
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func (b *Block) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("\n---\nTimestamp %x", b.Timestamp))
	for i, tx := range b.Transactions {
		lines = append(lines, fmt.Sprintf("--- Transaction %d", i))
		lines = append(lines, fmt.Sprintf("%v", tx))
	}
	lines = append(lines, fmt.Sprintf("--- PrevBlockHash %x", b.PrevBlockHash))
	lines = append(lines, fmt.Sprintf("--- Hash %x", b.Hash))
	lines = append(lines, fmt.Sprintf("--- Nonce %x", b.Nonce))
	lines = append(lines, fmt.Sprintf("--- Height %x", b.Height))
	return strings.Join(lines, "\n")
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Reach End of File")
			return &block
		}
		log.Panic(err)
	}
	return &block
}
