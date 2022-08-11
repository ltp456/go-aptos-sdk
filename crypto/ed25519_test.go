package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewKeyPairFromSeed(t *testing.T) {
	//seed, err := NewRandSeed()
	//if err != nil {
	//	panic(err)
	//}
	seed, _ := hex.DecodeString("76bcf7c263ab58224fc3f57d701de3836581df7d62b270a86344c5214812e0b9")
	fmt.Printf("seed:%x\n", seed)
	keyPair, err := NewKeyPairFromSeed(seed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private:%x\n", keyPair.Private())
	fmt.Printf("public:%x\n", keyPair.Public())
	fmt.Printf("address:%x\n", keyPair.Address())

	signature, err := keyPair.Sign(seed)
	if err != nil {
		panic(err)
	}

	fmt.Printf("signature: %x\n", signature)

	verify, err := keyPair.publickey.Verify(seed, signature)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify: %v\n", verify)

}
