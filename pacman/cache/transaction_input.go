package decentralizestorage

import (
	"bytes"
	cipher "home/pacman/cipher"
)

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := cipher.HashPubKey(in.PubKey)
	return bytes.Equal(lockingHash, pubKeyHash)
}
