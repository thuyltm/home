package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 10

// Vin: output of a or multiple previous transaction
// Vout: where coins are actually stored
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// Serialize returns a serialized Transaction
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte
	txCopy := *tx
	txCopy.ID = []byte{}
	hash = sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

// NewCoinbaseTX creates a first coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	// A coinbase transaction has only one input.
	// Txid is empty and Vout equal to -1.
	// Also, a coinbase transaction doesn't store a script in ScriptSig
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{10, to} //claim 10 free coins for TO address
	// In Bitcoin, this number is not stored anywhere
	// and calculated based only the total number of blocks
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

/*
NewUTXOTransaction creates a new transaction

	We want to send some coins to someone else
	1. Find all unspent outputs and ensure that they store enough value
	2. For each found output an input inferencing, we create two outputs
	  2.1 One that's locked with the receiver address. This is the actual
	  transfering of coins to other address
	  2.2 One that's locked with the sender address.
*/
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	// 1. Find all unspent outputs and ensure that they store enough value
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	// Build a list of TXInputs inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}
		for _, out := range outs {
			// referencing a Vout in an TXInput means that coin is being spent
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}
	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TXInput reference a specific, spendatable index within a transaction's Vouts list
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// TXOutput store "coins" (notice the Value field)
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

/**
** Only those balance can be unlocked by the key we own
(currently we use user defined addresses)
**/
// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(address string) bool {
	return in.ScriptSig == address
}

// CanBeUnlockedWith checks if the output can be unlocked with the provided address
func (out *TXOutput) CanBeUnlockedWith(address string) bool {
	return out.ScriptPubKey == address
}
