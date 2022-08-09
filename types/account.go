package types

type Account struct {
	SequenceNumber    string `json:"sequence_number"`
	AuthenticationKey int64  `json:"authentication_key"`
}
