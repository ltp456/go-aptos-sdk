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
	fmt.Printf("seed:0x%x\n", seed)
	keyPair, err := NewKeyPairFromSeed(seed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private:0x%x\n", keyPair.Private())
	fmt.Printf("public:0x%x\n", keyPair.Public())
	fmt.Printf("address:0x%x\n", keyPair.Address())

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

func TestNewKeyPairFromPrivateKey(t *testing.T) {
	//priv, err := hex.DecodeString("5660781d3a7c08343af80170b9f64d4f878d257a29870d22461f839da856c05f")
	//if err != nil {
	//	panic(err)
	//}

	//keyPair := NewKeyPairFromPrivateKey(priv)
	//fmt.Printf("%x\n", keyPair.Private())
	//fmt.Printf("%x\n", keyPair.Public())
	//fmt.Printf("%x\n", keyPair.Address())

}
