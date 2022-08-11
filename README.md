# go-apots-sdk
apots rest api implementation of the go language version

[Apots Node Api Doc](https://aptos.dev/api/latest-api.html#/)

## example


KeyPair

    seed, err := NewRandSeed()
    if err != nil {
        panic(err)
    }
    fmt.Printf("seed:%x\n", seed)


    keyPair, err := NewKeyPairFromSeed(seed)
    if err != nil {
        panic(err)
    }

    fmt.Printf("private:%x\n", keyPair.Private())
    fmt.Printf("public:%x\n", keyPair.Public())
    fmt.Printf("address:%x\n", keyPair.Address())


    signature, err := keyPair.Sign(message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("signature: %x\n", signature)

	verify, err := keyPair.publickey.Verify(message, signature)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify: %v\n", verify)




ApotsClient

    endpoint := "https://fullnode.devnet.aptoslabs.com"
    client, err = NewApotsClient(endpoint)
    if err != nil {
        panic(err)
    }

	ledgerInformation, err := client.LedgerInformation()
	if err != nil {
		panic(err)
	}
	fmt.Println(ledgerInformation)

