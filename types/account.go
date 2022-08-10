package types

type AccountModules struct {
	Bytecode string           `json:"bytecode"`
	Abi      AccountModuleAbi `json:"abi"`
}

type ExposedFunctions struct {
	Name              string              `json:"name"`
	Visibility        string              `json:"visibility"`
	GenericTypeParams []GenericTypeParams `json:"generic_type_params"`
	Params            []string            `json:"params"`
	Return            []interface{}       `json:"return"`
}

type GenericTypeParams struct {
	Constraints []interface{} `json:"constraints"`
	IsPhantom   bool          `json:"is_phantom"`
}

type Fields struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Structs struct {
	Name              string              `json:"name"`
	IsNative          bool                `json:"is_native"`
	Abilities         []string            `json:"abilities"`
	GenericTypeParams []GenericTypeParams `json:"generic_type_params"`
	Fields            []Fields            `json:"fields"`
}

type AccountModuleAbi struct {
	Address          string             `json:"address"`
	Name             string             `json:"name"`
	Friends          []string           `json:"friends"`
	ExposedFunctions []ExposedFunctions `json:"exposed_functions"`
	Structs          []Structs          `json:"structs"`
}

type Account struct {
	SequenceNumber    string `json:"sequence_number"`
	AuthenticationKey int64  `json:"authentication_key"`
}
type AccountResource struct {
	Type string  `json:"type"`
	Data ResData `json:"data"`
}

type Coin struct {
	Value string `json:"value"`
}

type ResData struct {
	Coin Coin `json:"coin"`
}
