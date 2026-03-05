package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// "target bits" is the block header storing the difficulity
// at which the block was mined
const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// In the NewProofOfWork function, we initialize a big.Int with the value of 1
// and shift it left by 256 - targetBits bits. 256 is the length of a SHA-256
// hash in bits, and it's SHA-256 hashing algorithm that we're going to use
// The hexadecimal representation of target is
// 0x10000000000000000000000000000000000000000000000000000000000
// And it occupies 29 bytes in memory
// You can think of a target as the upper boundary of a range: if a number (a hash)
// is lower than the boundary, it's valid, and vice versa
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

// We need the data to hash
// This data is straightforward: we just merge block fields with the target
// and nonce. nonce here is the counter from the Hashcash
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	//nonce is the counter
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		//1. Prepare data
		//2. Hash it with SHA-256
		//3. Convert the hash to a big integer
		//4. Compare the integer with the target
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}
