# bitshares

A Bitshares API consuming a websocket connection to an active full node or a RPC connection to your `cli_wallet`. 
Look for several examples in [examples](/examples) and [tests](/tests) folder. This is work in progress. To mitigate breaking changes, please use tagged branches. New tagged branches will be created for breaking changes. No additional cgo dependencies for transaction signing required. Use it at your own risk. 

## install

```bash
go get -u github.com/youthonline/bitshares
```

Install dev-dependencies with

```bash
make init
```

This API uses [ffjson](https://github.com/pquerna/ffjson). 
If you change types or operations you have to regenerate the required static `MarshalJSON` and `UnmarshalJSON` functions for the new/changed code.

```bash
make generate
```

If you encounter any errors, try: 

```bash
make generate_new
```

to generate ffjson helpers from scratch.

## generate operation samples
To generate op samples for testing, go to [gen](/gen) package.
Generated operation samples get injected automatically while running operation tests.

## testing
To test this stuff I use a combined docker based MainNet/TestNet wallet, you can find [here](https://github.com/youthonline/bitshares-docker).
Operations testing uses generated real blockchain sample code by [gen](/gen) package. To test run:

```bash
make test_operations
make test_api
```

or a long running block (deserialize/serialize/compare) range test.

```bash
make test_blocks
```

## code

```go
wsFullApiUrl := "wss://bitshares.openledger.info/ws"

api := bitshares.NewWebsocketAPI(wsFullApiUrl)
if err := api.Connect(); err != nil {
	log.Fatal(err)
}

api.OnError(func(err error) {
	log.Fatal(err)
})

UserID   := types.NewAccountID("1.2.253")
AssetBTS := types.NewAssetID("1.3.0")

res, api.GetAccountBalances(UserID, AssetBTS)
if err != nil {
	log.Fatal(err)
}

log.Printf("balances: %v", res)
```

If you need wallet functions, use:

```go
rpcApiUrl    := "http://localhost:8095" 
api := bitshares.NewWalletAPI(rpcApiUrl)

if err := api.Connect(); err != nil{
	log.Fatal(err)
}

...
```

For a long application lifecycle, you can use an API instance with latency tester that connects to the most reliable node.
Note: Because the tester takes time to unleash its magic, use the above-mentioned constructor for quick in and out.

```go
wsFullApiUrl := "wss://bitshares.openledger.info/ws"

//wsFullApiUrl serves as "quick startup" fallback endpoint here, 
//until the latency tester provides the first results.

api, err := bitshares.NewWithAutoEndpoint(wsFullApiUrl)
if err != nil {
	log.Fatal(err)
}

if err := api.Connect(); err != nil {
	log.Fatal(err)
}

api.OnError(func(err error) {
	log.Fatal(err)
})

...
```

## implemented and tested (serialize/unserialize) operations

- [x] OperationTypeTransfer OperationType
- [x] OperationTypeLimitOrderCreate
- [x] OperationTypeLimitOrderCancel
- [x] OperationTypeCallOrderUpdate
- [x] OperationTypeFillOrder (virtual)
- [x] OperationTypeAccountCreate
- [x] OperationTypeAccountUpdate
- [x] OperationTypeAccountWhitelist
- [x] OperationTypeAccountUpgrade
- [x] OperationTypeAccountTransfer
- [x] OperationTypeAssetCreate
- [x] OperationTypeAssetUpdate
- [x] OperationTypeAssetUpdateBitasset
- [x] OperationTypeAssetUpdateFeedProducers
- [x] OperationTypeAssetIssue
- [x] OperationTypeAssetReserve
- [x] OperationTypeAssetFundFeePool
- [x] OperationTypeAssetSettle
- [x] OperationTypeAssetGlobalSettle
- [x] OperationTypeAssetPublishFeed
- [x] OperationTypeWitnessCreate
- [x] OperationTypeWitnessUpdate
- [x] OperationTypeProposalCreate
- [x] OperationTypeProposalUpdate
- [x] OperationTypeProposalDelete
- [x] OperationTypeWithdrawPermissionCreate
- [x] OperationTypeWithdrawPermissionUpdate
- [x] OperationTypeWithdrawPermissionClaim
- [x] OperationTypeWithdrawPermissionDelete
- [x] OperationTypeCommitteeMemberCreate
- [x] OperationTypeCommitteeMemberUpdate
- [x] OperationTypeCommitteeMemberUpdateGlobalParameters
- [x] OperationTypeVestingBalanceCreate
- [x] OperationTypeVestingBalanceWithdraw
- [x] OperationTypeWorkerCreate
- [x] OperationTypeCustom
- [ ] OperationTypeAssert
- [x] OperationTypeBalanceClaim
- [x] OperationTypeOverrideTransfer
- [x] OperationTypeTransferToBlind
- [ ] OperationTypeBlindTransfer
- [x] OperationTypeTransferFromBlind
- [ ] OperationTypeAssetSettleCancel
- [x] OperationTypeAssetClaimFees
- [ ] OperationTypeFBADistribute
- [x] OperationTypeBidColatteral
- [ ] OperationTypeExecuteBid

## todo
- add missing operations
- add convenience functions 


Have fun and feel free to contribute needed operations and tests.