package decentralizestorage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	cipher "home/pacman/cipher"
	"log"
	"strings"
)

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// the output is locked to  a PubKeyHash, establishing ownership of the value for that key
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := cipher.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// Hashing is a one-way process, meaning it is irreversible
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

func (out *TXOutput) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("Value %d:", out.Value))
	lines = append(lines, fmt.Sprintf("PubKeyHash %x:", out.PubKeyHash))
	return strings.Join(lines, "\n")
}

func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))
	return txo
}

type TXOutputs struct {
	Outputs []TXOutput
}

func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (outs TXOutputs) String() string {
	var lines []string
	for i, out := range outs.Outputs {
		lines = append(lines, fmt.Sprintf("\n==================\nOut %d", i))
		lines = append(lines, fmt.Sprintf("Value %d", out.Value))
		lines = append(lines, fmt.Sprintf("PubKeyHash %x", out.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}

func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}
	return outputs
}
