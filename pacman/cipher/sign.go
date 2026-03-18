package cipher

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func Sign(privKey ecdsa.PrivateKey, transactionId []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, transactionId)
	if err != nil {
		return nil, err
	}
	return append(r.Bytes(), s.Bytes()...), nil
}

func Verify(transactionSignature []byte, senderPubKey []byte, transactionId []byte) bool {
	curve := elliptic.P256()
	r := big.Int{}
	s := big.Int{}
	sigLen := len(transactionSignature)
	r.SetBytes(transactionSignature[:(sigLen / 2)])
	s.SetBytes(transactionSignature[(sigLen / 2):])
	x := big.Int{}
	y := big.Int{}
	keyLen := len(senderPubKey)
	x.SetBytes(senderPubKey[:(keyLen / 2)])
	y.SetBytes(senderPubKey[(keyLen / 2):])
	rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}
	if ecdsa.Verify(&rawPubKey, transactionId, &r, &s) == false {
		return false
	}
	return true
}
