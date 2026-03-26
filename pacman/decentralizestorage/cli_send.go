package main

import (
	"fmt"
	. "home/pacman/cache"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
	bc := NewNodeBlockchain(nodeID)
	UTXOSet := UTXOSet{Blockchain: bc}
	defer bc.GetDB().Close()

	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)
	tx := NewNodeUTXOTransaction(&wallet, to, amount, &UTXOSet)
	if mineNow {
		txs := []*Transaction{tx}
		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}
	fmt.Println("\nSuccess!")
}
