package cli

import (
	"fmt"
	. "home/pacman/cache"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"

func (cli *CLI) printChain() {
	bc := NewBlockchain()
	defer bc.GetDB().Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) printBlockchain() {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			blockData := DeserializeBlock(v)
			fmt.Printf("Key: %x\n", k)
			fmt.Printf("Value: %+v\n", blockData)
		}
		return nil
	})

}

func (cli *CLI) printUTXO() {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("chainstate"))
		b.ForEach(func(k, v []byte) error {
			outs := DeserializeOutputs(v)
			fmt.Printf("Key: %x\n", k)
			fmt.Printf("Value: %v\n", outs)
			return nil
		})
		return nil
	})

}
