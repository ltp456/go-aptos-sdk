package types

type LedgerInformation struct {
	ChainID         int    `json:"chain_id"`
	LedgerVersion   string `json:"ledger_version"`
	LedgerTimestamp string `json:"ledger_timestamp"`
}
