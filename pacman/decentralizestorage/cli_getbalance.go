package main

import (
	"fmt"
	. "home/pacman/cache"
	. "home/pacman/cipher"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := NewNodeBlockchain(nodeID)
	UTXOSet := UTXOSet{Blockchain: bc}
	defer bc.GetDB().Close()

	balance := 0
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
