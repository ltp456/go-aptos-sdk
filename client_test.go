package go_apots_sdk

import (
	"fmt"
	"go-apots-sdk/form"
	"testing"
)

var client *ApotsClient
var err error

func init() {
	//endpoint := "http://aptos.dev/"
	endpoint := "https://fullnode.devnet.aptoslabs.com"
	client, err = NewApotsClient(endpoint)
	if err != nil {
		panic(err)
	}
}

func TestApotsClient_NewAddress(t *testing.T) {

	seed, err := client.NewRandSeed()
	if err != nil {
		panic(err)
	}

	keyPair, err := client.NewAccountFromSeed(seed)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("%x \n", keyPair.Private()))
	fmt.Println(fmt.Printf("%x \n", keyPair.Public()))
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

func TestApotsClient_Transactions(t *testing.T) {
	transactions, err := client.Transactions(0, 100)
	if err != nil {
		panic(err)
	}
	for _, tx := range transactions {
		fmt.Println(tx)
	}

}

func TestApotsClient_Transaction(t *testing.T) {
	transaction, err := client.Transaction("0x3044cfa1e88323298d8791e2539d192562e68f8640f9870149449becaf50e371")
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
	transaction, err := client.SimulateTransaction(
		"",
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
		"",
		form.Payload{},
		form.Signature{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(txSignMessage)
}

// ----account ------

func TestApotsClient_GetAccount(t *testing.T) {
	account, err := client.GetAccount("")
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}

func TestApotsClient_AccountResource(t *testing.T) {
	accountResources, err := client.AccountResource("", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(accountResources)
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
