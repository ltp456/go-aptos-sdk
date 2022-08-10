package go_apots_sdk

import (
	"fmt"
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

// ----- general -----

// ----account ------

func TestApotsClient_GetAccount(t *testing.T) {
	account, err := client.GetAccount("")
	if err != nil {
		panic(err)
	}
	fmt.Println(account)
}
