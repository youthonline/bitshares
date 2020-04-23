package tests

import (
	"testing"
	"time"

	"github.com/youthonline/bitshares"
	"github.com/youthonline/bitshares/config"
	"github.com/youthonline/bitshares/types"
	"github.com/stretchr/testify/suite"

	//import operations to initialize types.OperationMap
	_ "github.com/youthonline/bitshares/operations"
)

type commonTest struct {
	suite.Suite
	TestAPI bitshares.WebsocketAPI
}

func (suite *commonTest) SetupTest() {
	suite.TestAPI = NewWebsocketTestAPI(
		suite.T(),
		WsFullApiUrl,
	)
}

func (suite *commonTest) TearDownTest() {
	if err := suite.TestAPI.Close(); err != nil {
		suite.FailNow(err.Error(), "Close")
	}
}

func (suite *commonTest) Test_GetChainID() {
	res, err := suite.TestAPI.GetChainID()
	if err != nil {
		suite.FailNow(err.Error(), "GetChainID")
	}

	suite.Equal(res, config.ChainIDBTS)
}

func (suite *commonTest) Test_GetAccountBalances() {
	res, err := suite.TestAPI.GetAccountBalances(UserID2, AssetBTS)
	if err != nil {
		suite.FailNow(err.Error(), "GetAccountBalances 1")
	}

	suite.NotNil(res)
	//logging.Dump("balance bts >", res)

	res, err = suite.TestAPI.GetAccountBalances(UserID2)
	if err != nil {
		suite.FailNow(err.Error(), "GetAccountBalances 2")
	}

	suite.NotNil(res)
	//logging.Dump("balances all >", res)
}

func (suite *commonTest) Test_GetAccounts() {
	res, err := suite.TestAPI.GetAccounts(UserID2, UserID3, UserID4)
	if err != nil {
		suite.FailNow(err.Error(), "GetAccounts")
	}

	suite.NotNil(res)
	suite.Len(res, 3)

	//logging.Dump("get accounts >", res)
}

func (suite *commonTest) Test_GetFullAccounts() {
	res, err := suite.TestAPI.GetFullAccounts(UserID2, UserID3, UserID4)
	if err != nil {
		suite.FailNow(err.Error(), "GetFullAccounts")
	}

	suite.NotNil(res)
	suite.Len(res, 3)

	//logging.Dump("get full accounts >", res)
}

func (suite *commonTest) Test_GetTicker() {
	res, err := suite.TestAPI.GetTicker(AssetCNY, AssetBTS)
	if err != nil {
		suite.FailNow(err.Error(), "GetTicker")
	}

	suite.NotNil(res)
	//logging.Dump("get_ticker >", res)
}

func (suite *commonTest) Test_GetObjects() {
	res, err := suite.TestAPI.GetObjects(
		UserID1,
		AssetCNY,
		BitAssetDataCNY,
		// LimitOrder1,
		// CallOrder1,
		// SettleOrder1,
		OperationHistory1,
		CommitteeMember1,
		Balance1,
	)

	if err != nil {
		suite.FailNow(err.Error(), "GetObjects")
	}

	suite.NotNil(res)
	suite.Len(res, 6)
	//logging.Dump("objects >", res)
}

func (suite *commonTest) Test_GetBlock() {
	res, err := suite.TestAPI.GetBlock(33217575)
	if err != nil {
		suite.FailNow(err.Error(), "GetBlock")
	}

	suite.NotNil(res)
	//logging.Dump("get_block >", res)
}

func (suite *commonTest) Test_GetBlockHeader() {
	res, err := suite.TestAPI.GetBlockHeader(33217575)
	if err != nil {
		suite.FailNow(err.Error(), "GetBlockHeader")
	}

	suite.NotNil(res)
	//logging.Dump("get_block_header >", res)
}

func (suite *commonTest) Test_GetTransaction() {
	res, err := suite.TestAPI.GetTransaction(33217575, 1)
	if err != nil {
		suite.FailNow(err.Error(), "GetTransaction")
	}

	suite.NotNil(res)
	suite.Equal(res.RefBlockNum, types.UInt16(56357))
	//logging.Dump("get_transaction >", res)
}

func (suite *commonTest) Test_GetDynamicGlobalProperties() {
	res, err := suite.TestAPI.GetDynamicGlobalProperties()
	if err != nil {
		suite.FailNow(err.Error(), "GetDynamicGlobalProperties")
	}

	suite.NotNil(res)
	//logging.Dump("dynamic global properties >", res)
}

func (suite *commonTest) Test_GetAccountByName() {
	res, err := suite.TestAPI.GetAccountByName("openledger")
	if err != nil {
		suite.FailNow(err.Error(), "GetAccountByName")
	}

	suite.NotNil(res)
	//logging.Dump("accounts >", res)
}

func (suite *commonTest) Test_GetTradeHistory() {
	dtTo := time.Now().UTC()

	dtFrom := dtTo.Add(-time.Hour * 24)
	res, err := suite.TestAPI.GetTradeHistory(AssetBTS, AssetHERO, dtTo, dtFrom, 50)

	if err != nil {
		suite.FailNow(err.Error(), "GetTradeHistory")
	}

	suite.NotNil(res)
	//logging.Dump("markettrades >", res)
}

func (suite *commonTest) Test_GetLimitOrders() {
	res, err := suite.TestAPI.GetLimitOrders(AssetCNY, AssetBTS, 50)
	if err != nil {
		suite.FailNow(err.Error(), "GetLimitOrders")
	}

	suite.NotNil(res)
	//logging.Dump("limitorders >", res)
}

func (suite *commonTest) Test_GetCallOrders() {
	res, err := suite.TestAPI.GetCallOrders(AssetUSD, 50)
	if err != nil {
		suite.FailNow(err.Error(), "GetCallOrders")
	}

	suite.NotNil(res)
	//	logging.Dump("callorders >", res)
}

func (suite *commonTest) Test_GetMarginPositions() {
	res, err := suite.TestAPI.GetMarginPositions(UserID2)
	if err != nil {
		suite.FailNow(err.Error(), "GetMarginPositions")
	}

	suite.NotNil(res)
	//logging.Dump("marginpositions >", res)
}

func (suite *commonTest) Test_GetForceSettlementOrders() {
	res, err := suite.TestAPI.GetForceSettlementOrders(AssetCNY, 50)
	if err != nil {
		suite.FailNow(err.Error(), "GetSettleOrders")
	}

	suite.NotNil(res)
	//logging.Dump("settleorders >", res)logging.SetDebug(true)
}

func (suite *commonTest) Test_ListAssets() {
	res, err := suite.TestAPI.ListAssets("OPEN.DASH", 2)
	if err != nil {
		suite.FailNow(err.Error(), "ListAssets")
	}

	suite.NotNil(res)
	suite.Len(res, 2)
	//logging.Dump("assets >", res)
}

func (suite *commonTest) Test_GetAccountHistory() {

	user := types.NewAccountID("1.2.96393")
	start := types.NewOperationHistoryID("1.11.187698971")
	stop := types.NewOperationHistoryID("1.11.187658388")

	res, err := suite.TestAPI.GetAccountHistory(user, stop, 30, start)
	if err != nil {
		suite.FailNow(err.Error(), "GetAccountHistory")
	}

	suite.NotNil(res)
	//logging.Dump("history >", res)
}

//for qizikd
func (suite *commonTest) Test_GetAccountHistory_1() {

	user := types.NewAccountID("1.2.1587421")
	start := types.NewOperationHistoryID("1.11.0")
	stop := types.NewOperationHistoryID("1.11.187658388")

	res, err := suite.TestAPI.GetAccountHistory(user, stop, 30, start)
	if err != nil {
		suite.FailNow(err.Error(), "GetAccountHistory")
	}

	suite.NotNil(res)
	//logging.Dump("history >", res)
}

func (suite *commonTest) Test_GetOrderBook() {
	res, err := suite.TestAPI.GetOrderBook(AssetUSD, AssetBTS, 10)
	if err != nil {
		suite.FailNow(err.Error(), "GetOrderBook")
	}

	suite.NotNil(res)
	suite.Len(res.Asks, 10)
	suite.Len(res.Bids, 10)

	//logging.Dump("test GetOrderBook >", res)
}

func (suite *commonTest) Test_Get24Volume() {
	res, err := suite.TestAPI.Get24Volume(AssetUSD, AssetBTS)
	if err != nil {
		suite.FailNow(err.Error(), "Get24Volume")
	}

	suite.NotNil(res)
	suite.Equal("USD", res.Base.String())
	suite.Equal("BTS", res.Quote.String())

	//logging.Dump("test Get24Volume >", res)
}

func TestCommon(t *testing.T) {
	testSuite := new(commonTest)
	suite.Run(t, testSuite)
}
