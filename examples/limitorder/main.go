package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/youthonline/bitshares"
	"github.com/youthonline/bitshares/config"
	"github.com/youthonline/bitshares/crypto"
	"github.com/youthonline/bitshares/operations"
	"github.com/youthonline/bitshares/types"
	"github.com/juju/errors"
)

var (
	bts    = types.NewAssetID("1.3.0")
	cny    = types.NewAssetID("1.3.113")
	keyBag *crypto.KeyBag
	seller types.GrapheneObject
)

const (
	wsFullApiUrl = "wss://bitshares.openledger.info/ws"
)

func init() {
	// init is called before the API is initialized,
	// hence must define current chain config explicitly.
	config.SetCurrent(config.ChainIDBTS)

	sellerAcctID := os.Getenv("BTS_TEST_ACCOUNT")
	if sellerAcctID == "" {
		log.Fatalf("please set env var %q", "BTS_TEST_ACCOUNT")
	}
	seller = types.NewAccountID(sellerAcctID)

	testWIF := os.Getenv("BTS_TEST_WIF")
	if sellerAcctID == "" {
		log.Fatalf("please set env var %q", "BTS_TEST_WIF")
	}

	keyBag = crypto.NewKeyBag()
	if err := keyBag.Add(testWIF); err != nil {
		log.Fatal(errors.Annotate(err, "Add [wif]"))
	}
}

func main() {
	api := bitshares.NewWebsocketAPI(wsFullApiUrl)
	if err := api.Connect(); err != nil {
		log.Fatal(errors.Annotate(err, "OnConnect"))
	}

	api.OnError(func(err error) {
		log.Fatal(errors.Annotate(err, "OnError"))
	})

	op := operations.LimitOrderCreateOperation{
		FillOrKill: false,
		Seller:     types.AccountIDFromObject(seller),
		Extensions: types.Extensions{},
		AmountToSell: types.AssetAmount{
			Amount: 100,
			Asset:  types.AssetIDFromObject(bts),
		},
		MinToReceive: types.AssetAmount{
			Amount: 1000000,
			Asset:  types.AssetIDFromObject(cny),
		},
	}

	op.Expiration.Set(24 * time.Hour)
	tx, err := api.BuildSignedTransaction(keyBag, bts, &op)
	if err != nil {
		log.Fatal(errors.Annotate(err, "BuildSignedTransaction"))
	}

	if err := api.BroadcastTransaction(tx); err != nil {
		log.Fatal(errors.Annotate(err, "BroadcastTransaction"))
	}

	fmt.Println("operation successful broadcasted")
}
