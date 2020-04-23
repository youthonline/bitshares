package operations

//go:generate ffjson $GOFILE

import (
	"github.com/youthonline/bitshares/types"
	"github.com/youthonline/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetFundFeePool] = func() types.Operation {
		op := &AssetFundFeePoolOperation{}
		return op
	}
}

type AssetFundFeePoolOperation struct {
	types.OperationFee
	Amount      types.UInt64     `json:"amount"`
	AssetID     types.AssetID    `json:"asset_id"`
	Extensions  types.Extensions `json:"extensions"`
	FromAccount types.AccountID  `json:"from_account"`
}

func (p AssetFundFeePoolOperation) Type() types.OperationType {
	return types.OperationTypeAssetFundFeePool
}

func (p AssetFundFeePoolOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FromAccount); err != nil {
		return errors.Annotate(err, "encode new options")
	}

	if err := enc.Encode(p.AssetID); err != nil {
		return errors.Annotate(err, "encode asset id")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
