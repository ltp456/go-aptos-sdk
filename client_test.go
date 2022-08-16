package go_apots_sdk

import (
	"encoding/hex"
	"fmt"
	"github.com/ltp456/go-apots-sdk/form"
	"testing"
)

var client *ApotsClient
var err error

func init() {
	//endpoint := "http://aptos.dev/"
	//endpoint := "https://fullnode.devnet.aptoslabs.com"
	endpoint := "http://127.0.0.1:8080"
	client, err = NewApotsClient(endpoint)
	if err != nil {
		panic(err)
	}
}

// faucet

func TestApotsClient_FaucetFundAccount(t *testing.T) {
	//var faucetEndpoint = "https://faucet.devnet.aptoslabs.com"
	faucetEndpoint := "http://127.0.0.1:8081"
	client.SetFaucetEndpoint(faucetEndpoint)
	resp, err := client.FaucetFundAccount("0x0a01c2f21269ac1795fcb0fc50d6982a1ac1e198c5f5617f828bc9df4644db9c", 50000)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// ----- general -----

func TestApotsClient_LedgerInformation(t *testing.T) {
	ledgerInformation, err := client.LedgerInformation()
	if err != nil {
		panic(err)
	}
	fmt.Println(ledgerInformation)
}

func TestApotsClient_ApiDocument(t *testing.T) {
	apiDocument, err := client.apiDocument()
	if err != nil {
		panic(err)
	}
	fmt.Println(apiDocument)
}

func TestApotsClient_OpenApiDocument(t *testing.T) {
	openApiDocument, err := client.openApiDocument()
	if err != nil {
		panic(err)
	}
	fmt.Println(openApiDocument)
}

// ------- transaction -------

func TestApotsClient_Transfer(t *testing.T) {
	seedStr := "76bcf7c263ab58224fc3f57d701de3836581df7d62b270a86344c5214812e0b9"
	seed, err := hex.DecodeString(seedStr)
	if err != nil {
		panic(err)
	}
	sender := "0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9"
	recipient := "0x0a01c2f21269ac1795fcb0fc50d6982a1ac1e198c5f5617f828bc9df4644db9c"

	submitTransactionResp, err := client.Transfer(seed, sender, recipient, "100", "1000", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println(submitTransactionResp.Hash)

}

func TestApotsClient_Transactions(t *testing.T) {

	//transactions, err := client.Transactions(9371617, 100)
	transactions, err := client.Transactions(9371617, 1)
	if err != nil {
		panic(err)
	}
	for _, tx := range transactions {
		fmt.Println(tx)
	}

}

func TestApotsClient_Transaction(t *testing.T) {
	transaction, err := client.Transaction("0xe53332192af7675144a3ca0a21e5ca106929b411eee242b3149a2c83e0b60fe3")
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction)
}

func TestApotsClient_SubmitTransaction(t *testing.T) {
	transaction, err := client.SubmitTransaction(
		"",
		"",
		"",
		"",
		"",
		form.Payload{},
		form.Signature{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction)

}

func TestApotsClient_SimulateTransaction(t *testing.T) {
	seedStr := "76bcf7c263ab58224fc3f57d701de3836581df7d62b270a86344c5214812e0b9"
	seed, err := hex.DecodeString(seedStr)
	if err != nil {
		panic(err)
	}
	sender := "0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9"
	recipient := "0x6ac297031be21d7d3b83e53f76aa803016c389cd4bcdd4d0928b7aaa80c6ff83"

	transaction, err := client.SimulateTransaction(seed, sender, recipient, "10", "1000", "1", Value{"gas_currency_code", "XUS"})
	if err != nil {
		panic(err)
	}
	fmt.Println(transaction.Hash)
}

func TestApotsClient_AccountTransactions(t *testing.T) {
	transactions, err := client.AccountTransactions("", 0, 10)
	if err != nil {
		panic(err)
	}
	for _, tx := range transactions {
		fmt.Println(tx)
	}

}

func TestApotsClient_CreateTxSignMessage(t *testing.T) {
	txSignMessage, err := client.CreateTxSignMessage(
		"",
		"",
		"",
		"",
		"",
		form.Payload{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(txSignMessage)
}

// ----account ------

func TestApotsClient_GetAccount(t *testing.T) {
	account, err := client.GetAccount("0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9")
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}

func TestApotsClient_ApotsBalance(t *testing.T) {

	balance, err := client.ApotsBalance("0x0a01c2f21269ac1795fcb0fc50d6982a1ac1e198c5f5617f828bc9df4644db9c")
	if err != nil {
		panic(err)
	}
	fmt.Println(balance.String())
}

func TestApotsClient_AccountResource(t *testing.T) {
	accountResources, err := client.AccountResource("0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9")
	if err != nil {
		panic(err)
	}
	for _, item := range accountResources {
		fmt.Println(item.Type, item.Data)
	}
}

func TestApotsClient_AccountResourceByResType(t *testing.T) {
	accountResource, err := client.AccountResourceByResType("", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(accountResource)

}

func TestApotsClient_AccountModule(t *testing.T) {
	accountModules, err := client.AccountModules("", "")
	if err != nil {
		panic(err)
	}
	for _, m := range accountModules {
		fmt.Println(m)
	}

}

func TestApotsClient_AccountModuleByID(t *testing.T) {
	modules, err := client.AccountModuleByID("", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(modules)
}

// ---- events ------

func TestApotsClient_EventsByKey(t *testing.T) {
	events, err := client.EventsByKey("")
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		fmt.Println(event)
	}
}

func TestApotsClient_EventsByEventHandle(t *testing.T) {
	events, err := client.EventsByEventHandle("", "", "", 0, 10)
	if err != nil {
		panic(err)
	}
	for _, event := range events {
		fmt.Println(event)
	}
}

// ----- state -------

func TestApotsClient_TableItemByHandleAndKey(t *testing.T) {
	err := client.TableItemByHandleAndKey("", "", "", "")
	if err != nil {
		panic(err)
	}
}
