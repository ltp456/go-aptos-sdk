# go-apots-sdk
apots rest api implementation of the go language version

[Apots Node Api Doc](https://aptos.dev/api/latest-api.html#/)

### example

address

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




client

    client, err = NewApotsClient("https://fullnode.devnet.aptoslabs.com")
	if err != nil {
		panic(err)
	}