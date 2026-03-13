package main

import (
	"fmt"
)

func main() {
	wallets := NewWallets()
	address1 := wallets.CreateWallet()
	fmt.Printf("Your new address1: %s\n", address1)
	address2 := wallets.CreateWallet()
	fmt.Printf("Your new address2: %s\n", address2)
	wallets.SaveToFile()
	address3 := wallets.CreateWallet()
	fmt.Printf("Your new address: %s\n", address3)
	wallets.SaveToFile()
	otherWalletsDb := NewWallets()
	otherWalletsDb.LoadFromFile()
	fmt.Print(otherWalletsDb.GetWallet(address1))
	fmt.Print(otherWalletsDb.GetWallet(address2))
	fmt.Print(otherWalletsDb.GetWallet(address3))
}
