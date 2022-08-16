package types

type VMStatus string

const (
	ExecutedSuccessfully VMStatus = "Executed successfully"
)

type EventType string

const (
	WithdrawEvent EventType = "0x1::coin::WithdrawEvent"
	DepositEvent  EventType = "0x1::coin::DepositEvent"
)

type ResType string

const (
	ApotsCoinRes ResType = "0x1::coin::CoinStore<0x1::aptos_coin::AptosCoin>"
)

type SignatureType string

const (
	Ed25519Signature SignatureType = "ed25519_signature"
)

type FunctionType string

const (
	ScriptFunctionPayload FunctionType = "script_function_payload"
)

type Function string

const (
	CoinTransfer Function = "0x1::coin::transfer"
)

type CoinType string

const (
	ApotsCoin CoinType = "0x1::aptos_coin::AptosCoin"
)
