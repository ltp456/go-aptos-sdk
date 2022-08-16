package types

type SignMessage struct {
	Message string `json:"message"`
}

type SubmitTransactionResp struct {
	Type                    string          `json:"type"`
	Hash                    string          `json:"hash"`
	Sender                  string          `json:"sender"`
	SequenceNumber          string          `json:"sequence_number"`
	MaxGasAmount            string          `json:"max_gas_amount"`
	GasUnitPrice            string          `json:"gas_unit_price"`
	GasCurrencyCode         string          `json:"gas_currency_code"`
	ExpirationTimestampSecs string          `json:"expiration_timestamp_secs"`
	Payload                 SubmitTxPayload `json:"payload"`
	Signature               Signature       `json:"signature"`
}

type SubmitTxPayload struct {
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

type Transaction struct {
	Type                string    `json:"type"`
	Events              []Events  `json:"events"`
	Payload             Payload   `json:"payload"`
	Version             string    `json:"version"`
	Hash                string    `json:"hash"`
	Sender              string    `json:"sender"`
	StateRootHash       string    `json:"state_root_hash"`
	EventRootHash       string    `json:"event_root_hash"`
	GasUsed             string    `json:"gas_used"`
	Success             bool      `json:"success"`
	VMStatus            VMStatus  `json:"vm_status"`
	AccumulatorRootHash string    `json:"accumulator_root_hash"`
	Changes             []Changes `json:"changes"`
	Signature           Signature `json:"signature"`
	SequenceNumber      string    `json:"sequence_number"`
}

type Data struct {
	Created string `json:"created"`
	RoleID  string `json:"role_id"`
	Amount  string `json:"amount"`
}

type Events struct {
	Key            string    `json:"key"`
	SequenceNumber string    `json:"sequence_number"`
	Type           EventType `json:"type"`
	Data           Data      `json:"data"`
}

type Abi struct {
	Name              string              `json:"name"`
	Visibility        string              `json:"visibility"`
	GenericTypeParams []GenericTypeParams `json:"generic_type_params"`
	Params            []string            `json:"params"`
	Return            []interface{}       `json:"return"`
}

type Code struct {
	Bytecode string `json:"bytecode"`
	Abi      Abi    `json:"abi"`
}

type Script struct {
	Code          Code     `json:"code"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []string `json:"arguments"`
}

type WriteSet struct {
	Type      string `json:"type"`
	ExecuteAs string `json:"execute_as"`
	Script    Script `json:"script"`
}

type Payload struct {
	Type          FunctionType  `json:"type"`
	Function      Function      `json:"function"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
	WriteSet      WriteSet      `json:"write_set"`
}

type Changes struct {
	Type         string `json:"type"`
	StateKeyHash string `json:"state_key_hash"`
	Address      string `json:"address"`
	Module       string `json:"module"`
}
