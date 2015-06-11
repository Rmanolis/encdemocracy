package main

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

func main() {

	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L3072N256); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params
	dsa.GenerateKey(privatekey, rand.Reader) // this generates a public & private key pair

	var pubkey dsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privatekey)

	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubkey)
	fmt.Println(len(fmt.Sprint(pubkey)))
	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is the message to be signed and verified!")
	signhash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)

	// Verify
	verifystatus := dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true

	// we add additional data to change the signhash
	io.WriteString(h, "This message is NOT to be signed and verified!")
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be false
}
