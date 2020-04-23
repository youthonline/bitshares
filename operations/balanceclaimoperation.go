package operations

//go:generate ffjson $GOFILE

import (
	"github.com/youthonline/bitshares/types"
	"github.com/youthonline/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeBalanceClaim] = func() types.Operation {
		op := &BalanceClaimOperation{}
		return op
	}
}

type BalanceClaimOperation struct {
	types.OperationFee
	BalanceToClaim   types.BalanceID   `json:"balance_to_claim"`
	BalanceOwnerKey  types.PublicKey   `json:"balance_owner_key"`
	DepositToAccount types.AccountID   `json:"deposit_to_account"`
	TotalClaimed     types.AssetAmount `json:"total_claimed"`
}

func (p BalanceClaimOperation) Type() types.OperationType {
	return types.OperationTypeBalanceClaim
}

func (p BalanceClaimOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	return nil
}

func (p BalanceClaimOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.DepositToAccount); err != nil {
		return errors.Annotate(err, "encode DepositToAccount")
	}

	if err := enc.Encode(p.BalanceToClaim); err != nil {
		return errors.Annotate(err, "encode BalanceToClaim")
	}

	if err := enc.Encode(p.BalanceOwnerKey); err != nil {
		return errors.Annotate(err, "encode BalanceOwnerKey")
	}

	if err := enc.Encode(p.TotalClaimed); err != nil {
		return errors.Annotate(err, "encode TotalClaimed")
	}

	return nil
}
