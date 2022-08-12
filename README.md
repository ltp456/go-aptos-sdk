# go-apots-sdk
apots rest api implementation of the go language version

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
        information, err := apotsClient.LedgerInformation()
        if err != nil {
            panic(err)
        }
        fmt.Println(information)
    }
