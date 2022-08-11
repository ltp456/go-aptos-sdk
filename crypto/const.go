package crypto

const SeedLength = 32
const PublicKeyLength = 32
const SignatureLength = 64
const PrivateKeyLength = 64

type KeyType string

const (
	Ed25519Type   KeyType = "ed25519"
	Sr25519Type   KeyType = "Sr25519Type"
	Secp256k1Type KeyType = "Secp256k1Type"
)
