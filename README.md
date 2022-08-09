# go-apots-sdk
apots rest api implementation of the go language version

[Apots Node Api Doc](https://aptos.dev/api/latest-api.html#/)

### example

    client, err = NewApotsClient("https://fullnode.devnet.aptoslabs.com")
	if err != nil {
		panic(err)
	}