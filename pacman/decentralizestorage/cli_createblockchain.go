package main

import (
	"fmt"
	. "home/pacman/cache"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) createNodeBlockchain(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateNodeBlockchain(address, nodeID)
	defer bc.GetDB().Close()

	UTXOSet := UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()
	fmt.Println("Done!")

}
