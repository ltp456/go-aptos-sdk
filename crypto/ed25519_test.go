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
	seed, _ := hex.DecodeString("f164a8dc0eaf5b4293dbe62b703acdf7f7023400bbcc2751553a1127a52e4b3e")
	fmt.Printf("seed:%x\n", seed)
	keyPair, err := NewKeyPairFromSeed(seed)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private:%x\n", keyPair.Private())
	fmt.Printf("public:%x\n", keyPair.Public())
	fmt.Printf("address:%x\n", keyPair.Address())

	/*
		f164a8dc0eaf5b4293dbe62b703acdf7f7023400bbcc2751553a1127a52e4b3e
		&f164a8dc0eaf5b4293dbe62b703acdf7f7023400bbcc2751553a1127a52e4b3e18dea2e9118de1efa0783ca9cccb1634022c7c0a170743650c59073c427d17a0
		&18dea2e9118de1efa0783ca9cccb1634022c7c0a170743650c59073c427d17a0
	*/
}
