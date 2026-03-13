package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"log"
	"os"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const walletFile = "wallet.dat"
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

type TempWallet struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload) //conver byte array into a human readable string
	return address
}

// hash public Key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256() //generate a private key using Elliptic Curve Cryptography
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

func (w Wallet) MarshalBinary() (data []byte, err error) {
	privateKeyByte, err := convertECDSAPrivateKey(&w.PrivateKey)
	if err != nil {
		return nil, err
	}
	tempWallet := &TempWallet{
		PrivateKey: privateKeyByte,
		PublicKey:  w.PublicKey,
	}
	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err = encoder.Encode(tempWallet)
	if err != nil {
		log.Panic(err)
	}
	return content.Bytes(), nil
}

func (w *Wallet) UnmarshalBinary(data []byte) error {
	content := bytes.NewBuffer(data)
	tempWallet := &TempWallet{}
	decoder := gob.NewDecoder(content)
	err := decoder.Decode(&tempWallet)
	if err != nil {
		log.Panic(err)
	}
	privateKey, err := convertByteToECDSAPrivateKey(tempWallet.PrivateKey)
	if err != nil {
		return err
	}
	w.PrivateKey = *privateKey
	w.PublicKey = tempWallet.PublicKey
	return nil
}

func convertECDSAPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	// Marshal the private key to DER format
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}
	// Encode the DER bytes into PEM format
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}
	//Write the PEM data to the file
	return pem.EncodeToMemory(pemBlock), nil
}

func convertByteToECDSAPrivateKey(keyBytes []byte) (*ecdsa.PrivateKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, os.ErrInvalid
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func saveECDSAPrivateKey(key *ecdsa.PrivateKey, filename string) error {
	// Marshal the private key to DER format
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	// Encode the DER bytes into PEM format
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}
	//Write the PEM data to the file
	pemBytes := pem.EncodeToMemory(pemBlock)
	return os.WriteFile(filename, pemBytes, 0600)
}

func readECDSAPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Decode the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, os.ErrInvalid
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
