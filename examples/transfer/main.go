package main

import (
	"fmt"
	"log"

	"github.com/youthonline/bitshares"
	"github.com/youthonline/bitshares/config"
	"github.com/youthonline/bitshares/crypto"
	"github.com/youthonline/bitshares/tests"
	"github.com/youthonline/bitshares/types"
	"github.com/juju/errors"
)

var (
	from    = tests.TestAccount1ID
	to      = tests.TestAccount2ID
	test    = tests.AssetTEST
	memoWIF = tests.TestAccount1PrivKeyActive
	memo    = "my super secret memo message"

	keyBag *crypto.KeyBag
)

const (
	wsTestApiUrl = "wss://node.testnet.bitshares.eu/ws"
)

func init() {
	// init is called before the API is initialized,
	// hence must define current chain config explicitly.
	config.SetCurrent(config.ChainIDTest)
	keyBag = crypto.NewKeyBag()
	if err := keyBag.Add(memoWIF); err != nil {
		log.Fatal(errors.Annotate(err, "Add [wif]"))
	}
}

func main() {
	api := bitshares.NewWebsocketAPI(wsTestApiUrl)
	if err := api.Connect(); err != nil {
		log.Fatal(errors.Annotate(err, "OnConnect"))
	}

	api.OnError(func(err error) {
		log.Fatal(errors.Annotate(err, "OnError"))
	})

	amt := types.AssetAmount{
		Asset:  types.AssetIDFromObject(test),
		Amount: 10000,
	}

	err := api.Transfer(keyBag, from, to, test, amt, memo)
	if err != nil {
		log.Fatal(errors.Annotate(err, "Transfer"))
	}

	fmt.Println("transfer successful")
}
