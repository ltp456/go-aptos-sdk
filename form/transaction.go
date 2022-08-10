package form

type SubmitTransaction struct {
	Sender                  string    `json:"sender"`
	SequenceNumber          string    `json:"sequence_number"`
	MaxGasAmount            string    `json:"max_gas_amount"`
	GasUnitPrice            string    `json:"gas_unit_price"`
	GasCurrencyCode         string    `json:"gas_currency_code"`
	ExpirationTimestampSecs string    `json:"expiration_timestamp_secs"`
	Payload                 Payload   `json:"payload"`
	Signature               Signature `json:"signature"`
}

type Payload struct {
	Type          string   `json:"type"`
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []string `json:"arguments"`
}

type Signature struct {
	Type      string `json:"type"`
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}
