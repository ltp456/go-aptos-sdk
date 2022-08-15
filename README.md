# go-apots-sdk
apots rest api implementation of the go language version, generate address, sign transaction and others

[Apots Node Api Doc](https://aptos.dev/api/latest-api.html#/)

## example

install

    go get github.com/ltp456/go-apots-sdk


use example

    package main

    import (
        "fmt"
        goapotssdk "github.com/ltp456/go-apots-sdk"
        "github.com/ltp456/go-apots-sdk/crypto"
    )
    
    func main() {
    
        // keypair
        seed, err := crypto.NewRandSeed()
        if err != nil {
            panic(err)
        }
        fmt.Printf("seed: %x\n", seed)
        keyPair, err := crypto.NewKeyPairFromSeed(seed)
        if err != nil {
            panic(err)
        }
    
        fmt.Printf("prinvateKey: %x\n", keyPair.Private())
        fmt.Printf("publicKey: %x\n", keyPair.Public())
        fmt.Printf("address: %x\n", keyPair.Address())
    
        // sign and verify
        message := []byte("apots")
        signature, err := keyPair.Sign(message)
        if err != nil {
            panic(err)
        }
        publicKey := keyPair.Public()
        verify, err := publicKey.Verify(message, signature)
        if err != nil {
            panic(err)
        }
        fmt.Printf("verify: %v\n", verify)
        

        // client
        endpoint := "https://fullnode.devnet.aptoslabs.com"
        apotsClient, err := goapotssdk.NewApotsClient(endpoint)
        if err != nil {
            panic(err)
        }


        //faucet
        faucetEndpoint := "https://faucet.devnet.aptoslabs.com"
        client.SetFaucetEndpoint(faucetEndpoint)
        resp, err := apotsClient.FaucetFundAccount("0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9", 50000)
        if err != nil {
            panic(err)
        }
        fmt.Println(resp)


        // transfer 
        seedStr := ""
        seed, err := hex.DecodeString(seedStr)
        if err != nil {
            panic(err)
        }
        sender := "0x468f5ade8a4cb5e426bad07ad8d808fb067160bc506eab8620520f8a5a4a08c9"
        recipient := "0x6ac297031be21d7d3b83e53f76aa803016c389cd4bcdd4d0928b7aaa80c6ff83"
    
        submitTransactionResp, err := apotsClient.Transfer(seed, sender, recipient, "1000", "1000", "1")
        if err != nil {
            panic(err)
        }
        fmt.Println(submitTransactionResp.Hash)

        // Ledger
        information, err := apotsClient.LedgerInformation()
        if err != nil {
            panic(err)
        }
        fmt.Println(information)

    }
