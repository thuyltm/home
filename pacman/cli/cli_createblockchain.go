package cli

import (
	"fmt"
	. "home/pacman/cache"
	. "home/pacman/wallet"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address)
	defer bc.GetDB().Close()
	UTXOSet := UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()
	fmt.Println("Done!")
}
