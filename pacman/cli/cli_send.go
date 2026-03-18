package cli

import (
	"fmt"
	. "home/pacman/cache"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain()
	UTXOSet := UTXOSet{Blockchain: bc}
	defer bc.GetDB().Close()

	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	txs := []*Transaction{tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")
}
