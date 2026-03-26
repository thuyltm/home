package main

import (
	"fmt"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()
	for _, address := range addresses {
		fmt.Println(address)
	}
}
