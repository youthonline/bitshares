// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package types

import "strconv"

const _OperationType_name = "OperationTypeTransferOperationTypeLimitOrderCreateOperationTypeLimitOrderCancelOperationTypeCallOrderUpdateOperationTypeFillOrderOperationTypeAccountCreateOperationTypeAccountUpdateOperationTypeAccountWhitelistOperationTypeAccountUpgradeOperationTypeAccountTransferOperationTypeAssetCreateOperationTypeAssetUpdateOperationTypeAssetUpdateBitassetOperationTypeAssetUpdateFeedProducersOperationTypeAssetIssueOperationTypeAssetReserveOperationTypeAssetFundFeePoolOperationTypeAssetSettleOperationTypeAssetGlobalSettleOperationTypeAssetPublishFeedOperationTypeWitnessCreateOperationTypeWitnessUpdateOperationTypeProposalCreateOperationTypeProposalUpdateOperationTypeProposalDeleteOperationTypeWithdrawPermissionCreateOperationTypeWithdrawPermissionUpdateOperationTypeWithdrawPermissionClaimOperationTypeWithdrawPermissionDeleteOperationTypeCommitteeMemberCreateOperationTypeCommitteeMemberUpdateOperationTypeCommitteeMemberUpdateGlobalParametersOperationTypeVestingBalanceCreateOperationTypeVestingBalanceWithdrawOperationTypeWorkerCreateOperationTypeCustomOperationTypeAssertOperationTypeBalanceClaimOperationTypeOverrideTransferOperationTypeTransferToBlindOperationTypeBlindTransferOperationTypeTransferFromBlindOperationTypeAssetSettleCancelOperationTypeAssetClaimFeesOperationTypeFBADistributeOperationTypeBidColatteralOperationTypeExecuteBid"

var _OperationType_index = [...]uint16{0, 21, 50, 79, 107, 129, 155, 181, 210, 237, 265, 289, 313, 345, 382, 405, 430, 459, 483, 513, 542, 568, 594, 621, 648, 675, 712, 749, 785, 822, 856, 890, 940, 973, 1008, 1033, 1052, 1071, 1096, 1125, 1153, 1179, 1209, 1239, 1266, 1292, 1318, 1341}

func (i OperationType) String() string {
	if i < 0 || i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
