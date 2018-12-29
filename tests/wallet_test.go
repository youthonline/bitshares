package tests

import (
	"log"
	"math/rand"
	"testing"

	"github.com/denkhaus/bitshares/api"
	"github.com/denkhaus/bitshares/config"
	"github.com/denkhaus/bitshares/crypto"
	"github.com/denkhaus/bitshares/operations"
	"github.com/denkhaus/bitshares/types"
	"github.com/juju/errors"
	"github.com/stretchr/testify/suite"
)

type walletAPITest struct {
	suite.Suite
	TestAPI    api.WalletAPI
	FullAPI    api.WalletAPI
	TestKeyBag *crypto.KeyBag
}

func (suite *walletAPITest) SetupSuite() {
	suite.FullAPI = NewWalletTestAPI(suite.T(), RpcFullApiUrl, config.ChainIDBTS)
	suite.TestAPI = NewWalletTestAPI(suite.T(), RpcTestApiUrl, config.ChainIDTest)

	suite.TestKeyBag = crypto.NewKeyBag()
	if err := suite.TestKeyBag.Add(TestAccount1PrivKeyActive); err != nil {
		suite.FailNow(err.Error(), "KeyBag.Add 1")
	}

	if err := suite.TestKeyBag.Add(TestAccount2PrivKeyActive); err != nil {
		suite.FailNow(err.Error(), "KeyBag.Add 2")
	}

	if err := suite.TestAPI.Unlock("123456"); err != nil {
		suite.FailNow(err.Error(), "Unlock")
	}
}

func (suite *walletAPITest) TearDownSuite() {
	if err := suite.TestAPI.Lock(); err != nil {
		suite.FailNow(err.Error(), "Lock")
	}

	if err := suite.TestAPI.Close(); err != nil {
		suite.FailNow(err.Error(), "Close [test api]")
	}

	if err := suite.FullAPI.Close(); err != nil {
		suite.FailNow(err.Error(), "Close [full api]")
	}
}

func (suite *walletAPITest) Test_GetBlock() {
	config.SetCurrentConfig(config.ChainIDBTS)
	block, err := suite.FullAPI.GetBlock(33451303)
	if err != nil {
		suite.FailNow(err.Error(), "GetBlock")
	}

	suite.Len(block.Transactions, 37)
	//logging.Dump("wallet_get_block <", block)
}

func (suite *walletAPITest) Test_GetDynamicGlobalProperties() {
	props, err := suite.FullAPI.GetDynamicGlobalProperties()
	if err != nil {
		suite.FailNow(err.Error(), "GetDynamicGlobalProperties")
	}

	suite.NotNil(props)
	suite.Equal(props.ID, *types.NewGrapheneID("2.1.0"))
	//logging.Dump("wallet_get_dynamic_global_properties <", props)
}

// func (suite *walletAPITest) Test_Transfer2() {
// 	trx, err := suite.TestAPI.Transfer2(
// 		TestAccount2ID,
// 		TestAccount1ID,
// 		"10",
// 		AssetTEST,
// 		"this is my transfer2",
// 	)

// 	if err != nil {
// 		suite.FailNow(err.Error(), "Transfer2")
// 	}

// 	logging.Dump("wallet_transfer2 <", trx)

// }

func (suite *walletAPITest) Test_GetRelativeAccountHistory() {
	hists, err := suite.TestAPI.GetRelativeAccountHistory(
		TestAccount1ID,
		0, 10, 0,
	)

	if err != nil {
		suite.FailNow(err.Error(), "GetRelativeAccountHistory")
	}

	suite.Len(hists, 10)
	//logging.Dump("wallet_get_relative_account_history <", hists)
}

func (suite *walletAPITest) Test_ReadMemo() {
	config.SetCurrentConfig(config.ChainIDTest)

	pubKeyA, err := types.NewPublicKeyFromString(TestAccount1PubKeyActive)
	if err != nil {
		log.Fatal(errors.Annotate(err, "NewPublicKeyFromString [key A]"))
	}
	pubKeyB, err := types.NewPublicKeyFromString(TestAccount2PubKeyActive)
	if err != nil {
		log.Fatal(errors.Annotate(err, "NewPublicKeyFromString [key B]"))
	}

	memo := types.Memo{
		From:  *pubKeyA,
		To:    *pubKeyB,
		Nonce: types.UInt64(rand.Int63()),
	}

	privKeyA, err := types.NewPrivateKeyFromWif(TestAccount1PrivKeyActive)
	if err != nil {
		log.Fatal(errors.Annotate(err, "NewPrivateKeyFromWif [key A]"))
	}

	memoSrc := "This is my exciting memo message"
	if err := memo.Encrypt(privKeyA, memoSrc); err != nil {
		log.Fatal(errors.Annotate(err, "Encrypt"))
	}

	msg, err := suite.TestAPI.ReadMemo(
		&memo,
	)

	if err != nil {
		suite.FailNow(err.Error(), "ReadMemo")
	}

	suite.Equal(memoSrc, msg)
}

func (suite *walletAPITest) Test_TransferExtended() {
	props, err := suite.TestAPI.GetDynamicGlobalProperties()
	if err != nil {
		suite.FailNow(err.Error(), "GetDynamicGlobalProperties")
	}

	trx, err := types.NewSignedTransactionWithBlockData(props)
	if err != nil {
		suite.FailNow(err.Error(), "NewSignedTransactionWithBlockData")
	}

	trx.Operations = types.Operations{
		&operations.TransferOperation{
			Extensions: types.Extensions{},
			Amount: types.AssetAmount{
				Amount: 100000,
				Asset:  *AssetTEST,
			},
			From: *TestAccount1ID,
			To:   *TestAccount2ID,
		},
	}

	// logging.SetDebug(true)
	// defer logging.SetDebug(false)

	fees, err := suite.TestAPI.GetRequiredFees(trx.Operations, AssetTEST)
	if err != nil {
		suite.FailNow(err.Error(), "GetRequiredFees")
	}

	if err := trx.Operations.ApplyFees(fees); err != nil {
		suite.FailNow(err.Error(), "ApplyFees")
	}

	suite.compareTransaction(trx, false)

	res, err := suite.TestAPI.SignTransaction(trx, true)
	if err != nil {
		suite.FailNow(err.Error(), "SignTransaction")
	}

	_ = res
	//logging.Dump("transfer <", res)
}

func (suite *walletAPITest) compareTransaction(tx *types.SignedTransaction, debug bool) {
	ref, test, err := CompareTransactions(suite.TestAPI, tx, debug)
	if err != nil {
		suite.FailNow(err.Error(), "Compare Transactions")
	}

	suite.Equal(ref, test)
}

func TestWalletAPI(t *testing.T) {
	testSuite := new(walletAPITest)
	suite.Run(t, testSuite)
}
