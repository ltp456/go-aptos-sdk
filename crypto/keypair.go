package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

type IKeypair interface {
	Type() KeyType
	Sign(msg []byte) ([]byte, error)
	Public() PublicKey
	Private() PrivateKey
	Address() []byte
}

// todo

type PrivateKey ed25519.PrivateKey
type PublicKey ed25519.PublicKey
type SecretKey []byte

func NewRandSeed() ([]byte, error) {
	buf := make([]byte, SeedLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
