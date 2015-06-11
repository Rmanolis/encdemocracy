package ecrypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"hash"
	"io"
	"log"
	"math/big"
)

func CreateKey() (*ecdsa.PrivateKey, error) {
	pubkeyCurve := elliptic.P521()

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return privatekey, nil
}

type Signature struct {
	R    *big.Int
	S    *big.Int
	Hash []byte
}

func GetSignature(pk *ecdsa.PrivateKey, msg string) (*Signature, error) {
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, msg)
	signhash := h.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, pk, signhash)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	sgn := new(Signature)
	sgn.R = r
	sgn.S = s
	sgn.Hash = signhash
	return sgn, nil
}
