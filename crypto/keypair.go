package crypto

import (
	"crypto/rand"
)

// todo maybe more type

type IKeypair interface {
	Type() KeyType
	Sign(msg []byte) ([]byte, error)
	Public() IPrivateKey
	Private() IPublicKey
	Address() []byte
	Verify(message, signature []byte) (bool, error)
}

type IPrivateKey interface {
}

type IPublicKey interface {
}

func NewRandSeed() ([]byte, error) {
	buf := make([]byte, SeedLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
