package crypto

import (
	"crypto/ed25519"
	"fmt"
	"golang.org/x/crypto/sha3"
)

type KeyPair struct {
	private   *PrivateKey
	publickey *PublicKey
	signKey   []byte
}

func (kp *KeyPair) Type() KeyType {
	return Ed25519Type
}

func (kp *KeyPair) Public() PublicKey {
	return *kp.publickey
}

func (kp *KeyPair) Private() PrivateKey {
	return *kp.private
}

func (kp *KeyPair) Address() []byte {
	hash := sha3.New256()
	hash.Write(kp.Public())
	bytes := make([]byte, 1)
	hash.Write(bytes)
	return hash.Sum(nil)
}

func (kp *KeyPair) Sign(msg []byte) ([]byte, error) {
	return nil, nil
}

func NewKeyPairFromPrivateKey(edPriv ed25519.PrivateKey) *KeyPair {
	publicKey := PublicKey(edPriv.Public().(ed25519.PublicKey))
	privateKey := PrivateKey(edPriv)
	keyPair := &KeyPair{
		publickey: &publicKey,
		private:   &privateKey,
	}
	return keyPair
}

func NewKeyPairFromSeed(seed []byte) (*KeyPair, error) {
	if len(seed) != SeedLength {
		return nil, fmt.Errorf("seed length must 32,  actually %v", len(seed))
	}
	edPriv := ed25519.NewKeyFromSeed(seed)

	publicKey := PublicKey(edPriv.Public().(ed25519.PublicKey))
	privateKey := PrivateKey(edPriv)
	keyPair := &KeyPair{
		publickey: &publicKey,
		private:   &privateKey,
	}
	return keyPair, nil
}
