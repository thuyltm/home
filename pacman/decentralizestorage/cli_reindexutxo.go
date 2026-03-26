package main

import (
	"fmt"
	. "home/pacman/cache"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := NewNodeBlockchain(nodeID)
	UTXOSet := UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
