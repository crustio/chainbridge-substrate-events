package events

import (
	"fmt"
	"math/big"

	"github.com/crustio/go-substrate-rpc-client/v4/scale"
	"github.com/crustio/go-substrate-rpc-client/v4/types"
)

type Events struct {
	ChainBridge_FungibleTransfer        []EventFungibleTransfer         //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer      //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer          //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged  //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted         //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded             //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved           //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                  //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst              //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved         //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected         //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded        //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed           //nolint:stylecheck,golint
	BridgeTransfer_FeeUpdated           []EventBridgeTransferFeeUpdated //nolint:stylecheck,golint
}

type CrustEvents struct {
	Staking_EraReward                     []EventEraReward                            //nolint:stylecheck,golint
	Staking_ValidateSuccess               []EventValidateSuccess                      //nolint:stylecheck,golint
	Staking_GuaranteeSuccess              []EventGuaranteeSuccess                     //nolint:stylecheck,golint
	Staking_CutGuaranteeSuccess           []EventCutGuaranteeSuccess                  //nolint:stylecheck,golint
	Staking_ChillSuccess                  []EventChillSuccess                         //nolint:stylecheck,golint
	Staking_UpdateStakeLimitSuccess       []EventUpdateStakeLimitSuccess              //nolint:stylecheck,golint
	Staking_UpdateIdentitiesSuccess       []EventUpdateIdentitiesSuccess              //nolint:stylecheck,golint
	Staking_PotList                       []EventPotList                              //nolint:stylecheck,golink
	Swork_RegisterSuccess                 []EventRegisterSuccess                      //nolint:stylecheck,golint
	Swork_WorksReportSuccess              []EventWorksReportSuccess                   //nolint:stylecheck,golint
	Swork_ABUpgradeSuccess                []EventABUpgradeSuccess                     //nolint:stylecheck,golint
	Swork_SetCodeSuccess                  []EventSetCodeSuccess                       //nolint:stylecheck,golint
	Swork_JoinGroupSuccess                []EventJoinGroupSuccess                     //nolint:stylecheck,golint
	Swork_QuitGroupSuccess                []EventQuitGroupSuccess                     //nolint:stylecheck,golint
	Swork_CreateGroupSuccess              []EventCreateGroupSuccess                   //nolint:stylecheck,golint
	Swork_KickOutSuccess                  []EventKickOutSuccess                       //nolint:stylecheck,golint
	Swork_AddIntoAllowlistSuccess         []EventAddIntoAllowlistSuccess              //nolint:stylecheck,golint
	Swork_RemoveFromAllowlistSuccess      []EventRemoveFromAllowlistSuccess           //nolint:stylecheck,golint
	Swork_SetPunishmentSuccess            []EventSetPunishmentSuccess                 //nolint:stylecheck,golint
	Swork_RemoveCodeSuccess               []EventRemoveCodeSuccess                    //nolint:stylecheck,golint
	Swork_ChillSuccess                    []EventSworkChillSuccess                    //nolint:stylecheck,golint
	Swork_CancelPunishmentSuccess         []EventCancelPunishmentSuccess              //nolint:stylecheck,golint
	Swork_QueueWorkReportSuccess          []EventQueueWorkReportSuccess               //nolint:stylecheck,golint
	Swork_SetSpowerSuperiorSuccess        []EventSetSpowerSuperiorSuccess             //nolint:stylecheck,golint
	Swork_UpdateSpowerSuccess             []EventUpdateSpowerSuccess                  //nolint:stylecheck,golint
	Market_FileSuccess                    []EventFileSuccess                          //nolint:stylecheck,golint
	Market_RenewFileSuccess               []EventRenewFileSuccess                     //nolint:stylecheck,golint
	Market_AddPrepaidSuccess              []EventAddPrepaidSuccess                    //nolint:stylecheck,golint
	Market_CalculateSuccess               []EventCalculateSuccess                     //nolint:stylecheck,golint
	Market_IllegalFileClosed              []EventIllegalFileClosed                    //nolint:stylecheck,golint
	Market_RewardMerchantSuccess          []EventRewardMerchantSuccess                //nolint:stylecheck,golint
	Market_SetEnableMarketSuccess         []EventSetEnableMarketSuccess               //nolint:stylecheck,golint
	Market_SetBaseFeeSuccess              []EventSetBaseFeeSuccess                    //nolint:stylecheck,golint
	Market_RegisterSuccess                []EventMarketRegisterSuccess                //nolint:stylecheck,golint
	Market_AddCollateralSuccess           []EventAddCollateralSuccess                 //nolint:stylecheck,golint
	Market_CutCollateralSuccess           []EventCutCollateralSuccess                 //nolint:stylecheck,golint
	Market_PotList                        []EventMarketPotList                        //nolint:stylecheck,golint
	Market_SetMarketSwitchSuccess         []EventSetMarketSwitchSuccess               //nolint:stylecheck,golint
	Market_SetFreeOrderAdminSuccess       []EventSetFreeOrderAdminSuccess             //nolint:stylecheck,golint
	Market_NewFreeAccount                 []EventNewFreeAccount                       //nolint:stylecheck,golint
	Market_SetFreeFeeSuccess              []EventSetFreeFeeSuccess                    //nolint:stylecheck,golint
	Market_FreeAccountRemoved             []EventFreeAccountRemoved                   //nolint:stylecheck,golint
	Market_SetFreeCountsLimitSuccess      []EventSetFreeCountsLimitSuccess            //nolint:stylecheck,golint
	Market_SetTotalFreeFeeLimitSuccess    []EventSetTotalFreeFeeLimitSuccess          //nolint:stylecheck,golint
	Market_SetSpowerSuperiorSuccess       []EventSetSpowerSuperiorSuccess             //nolint:stylecheck,golint
	Market_UpdateReplicasSuccess          []EventUpdateReplicasSuccess                //nolint:stylecheck,golint
	Locks_UnlockStartedFrom               []EventUnlockStartedFrom                    //nolint:stylecheck,golint
	Locks_UnlockSuccess                   []EventUnlockSuccess                        //nolint:stylecheck,golint
	Claims_SuperiorChanged                []EventSuperiorChanged                      //nolint:stylecheck,golint
	Claims_MinerChanged                   []EventMinerChanged                         //nolint:stylecheck,golint
	Claims_SetLimitSuccess                []EventSetLimitSuccess                      //nolint:stylecheck,golint
	Claims_MintSuccess                    []EventMintSuccess                          //nolint:stylecheck,golint
	Claims_Claimed                        []EventClaimsClaimed                        //nolint:stylecheck,golint
	Claims_BondEthSuccess                 []EventBondEthSuccess                       //nolint:stylecheck,golint
	Claims_Cru18MinerChanged              []EventCru18MinerChanged                    //nolint:stylecheck,golint
	Claims_Cru18MintSuccess               []EventCru18MintSuccess                     //nolint:stylecheck,golint
	Claims_Cru18Claimed                   []EventCru18Claimed                         //nolint:stylecheck,golint
	Claims_CsmMinerChanged                []EventCsmMinerChanged                      //nolint:stylecheck,golint
	Claims_SetCsmLimitSuccess             []EventSetCsmLimitSuccess                   //nolint:stylecheck,golint
	Claims_CsmMintSuccess                 []EventCsmMintSuccess                       //nolint:stylecheck,golint
	Claims_CsmClaimed                     []EventCsmClaimed                           //nolint:stylecheck,golint
	Benefits_AddBenefitFundsSuccess       []EventAddBenefitFundsSuccess               //nolint:stylecheck,golint
	Benefits_CutBenefitFundsSuccess       []EventCutBenefitFundsSuccess               //nolint:stylecheck,golint
	Benefits_RebondBenefitFundsSuccess    []EventRebondBenefitFundsSuccess            //nolint:stylecheck,golint
	Benefits_WithdrawBenefitFundsSuccess  []EventWithdrawBenefitFundsSuccess          //nolint:stylecheck,golint
	CSMLocking_Bonded                     []EventCSMLockingBonded                     //nolint:stylecheck,golint
	CSMLocking_Unbonded                   []EventCSMLockingUnbonded                   //nolint:stylecheck,golint
	CSMLocking_Withdrawn                  []EventCSMLockingWithdrawn                  //nolint:stylecheck,golint
	CSMLocking_SetCSMGuaranteePerfSuccess []EventCSMLockingSetCSMGuaranteePerfSuccess //nolint:stylecheck,golint
	CSMLocking_CSMGuaranteeSuccess        []EventCSMLockingCSMGuaranteeSuccess        //nolint:stylecheck,golint
	CSMLocking_CancelCSMGuaranteeSuccess  []EventCSMLockingCancelCSMGuaranteeSuccess  //nolint:stylecheck,golint
	Candy_CandyIssued                     []EventCandyCandyIssued                     //nolint:stylecheck,golint
	Candy_CandyTransferred                []EventCandyCandyTransferred                //nolint:stylecheck,golint
	Candy_CandyBurned                     []EventCandyCandyBurned                     //nolint:stylecheck,golint
	CSM_Endowed                           []types.EventBalancesEndowed                //nolint:stylecheck,golint
	CSM_DustLost                          []types.EventBalancesDustLost               //nolint:stylecheck,golint
	CSM_Transfer                          []types.EventBalancesTransfer               //nolint:stylecheck,golint
	CSM_BalanceSet                        []types.EventBalancesBalanceSet             //nolint:stylecheck,golint
	CSM_Deposit                           []types.EventBalancesDeposit                //nolint:stylecheck,golint
	CSM_Reserved                          []types.EventBalancesReserved               //nolint:stylecheck,golint
	CSM_Unreserved                        []types.EventBalancesUnreserved             //nolint:stylecheck,golint
	CSM_ReservedRepatriated               []types.EventBalancesReserveRepatriated     //nolint:stylecheck,golint
}

type EventEraReward struct {
	Phase           types.Phase
	EraIndex        types.U32
	AuthoringReward types.U128
	StakingReward   types.U128
	Topics          []types.Hash
}

type EventValidateSuccess struct {
	Phase        types.Phase
	Who          types.AccountID
	GuaranteeFee Perbill
	Topics       []types.Hash
}

type PerbillValue types.U32

type Perbill struct {
	Value PerbillValue
}

func (d *Perbill) Decode(decoder scale.Decoder) error {
	_, err := decoder.DecodeUintCompact()
	return err
}

func (d Perbill) Encode(encoder scale.Encoder) error {
	encoder.EncodeUintCompact(*big.NewInt(0).SetUint64(uint64(d.Value)))
	return nil
}

type EventGuaranteeSuccess struct {
	Phase  types.Phase
	Source types.AccountID
	Target types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCutGuaranteeSuccess struct {
	Phase  types.Phase
	Source types.AccountID
	Target types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventChillSuccess struct {
	Phase      types.Phase
	Stash      types.AccountID
	Controller types.AccountID
	Topics     []types.Hash
}

type EventUpdateStakeLimitSuccess struct {
	Phase  types.Phase
	Number types.U32
	Topics []types.Hash
}

type EventUpdateIdentitiesSuccess struct {
	Phase    types.Phase
	EraIndex types.U32
	Topics   []types.Hash
}

type EventPotList struct {
	Phase   types.Phase
	Address types.AccountID
	Topics  []types.Hash
}

type EventRegisterSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	PubKey types.Bytes
	Topics []types.Hash
}

type EventWorksReportSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	PubKey types.Bytes
	Topics []types.Hash
}
type EventABUpgradeSuccess struct {
	Phase     types.Phase
	Who       types.AccountID
	OldPubKey types.Bytes
	NewPubKey types.Bytes
	Topics    []types.Hash
}

type EventSetCodeSuccess struct {
	Phase       types.Phase
	Code        types.Bytes
	BlockNumber types.U32
	Topics      []types.Hash
}

type EventJoinGroupSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Group  types.AccountID
	Topics []types.Hash
}

type EventQuitGroupSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Group  types.AccountID
	Topics []types.Hash
}

type EventCreateGroupSuccess struct {
	Phase  types.Phase
	Group  types.AccountID
	Topics []types.Hash
}

type EventKickOutSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventAddIntoAllowlistSuccess struct {
	Phase  types.Phase
	Group  types.AccountID
	Who    types.AccountID
	Topics []types.Hash
}

type EventRemoveFromAllowlistSuccess struct {
	Phase  types.Phase
	Group  types.AccountID
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetPunishmentSuccess struct {
	Phase        types.Phase
	IsPunishment bool
	Topics       []types.Hash
}

type EventRemoveCodeSuccess struct {
	Phase  types.Phase
	Code   types.Bytes
	Topics []types.Hash
}

type EventSworkChillSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	PubKey types.Bytes
	Topics []types.Hash
}

type EventCancelPunishmentSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventFileSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Cid    types.Bytes
	Topics []types.Hash
}

type EventRenewFileSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Cid    types.Bytes
	Topics []types.Hash
}

type EventAddPrepaidSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Cid    types.Bytes
	Amount types.U128
	Topics []types.Hash
}

type EventCalculateSuccess struct {
	Phase  types.Phase
	Cid    types.Bytes
	Topics []types.Hash
}

type EventIllegalFileClosed struct {
	Phase  types.Phase
	Cid    types.Bytes
	Topics []types.Hash
}

type EventRewardMerchantSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetEnableMarketSuccess struct {
	Phase    types.Phase
	IsEnable bool
	Topics   []types.Hash
}

type EventSetBaseFeeSuccess struct {
	Phase  types.Phase
	Amount types.U128
	Topics []types.Hash
}

type EventMarketRegisterSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventAddCollateralSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCutCollateralSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventMarketPotList struct {
	Phase             types.Phase
	CollateralAddress types.AccountID
	StorageAddress    types.AccountID
	StakingAddress    types.AccountID
	ReservedAddress   types.AccountID
	Topics            []types.Hash
}

type EventSetMarketSwitchSuccess struct {
	Phase    types.Phase
	IsEnable bool
	Topics   []types.Hash
}

type EventSetFreeOrderAdminSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventNewFreeAccount struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetFreeFeeSuccess struct {
	Phase  types.Phase
	Amount types.U128
	Topics []types.Hash
}

type EventFreeAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetFreeCountsLimitSuccess struct {
	Phase  types.Phase
	Count  types.U32
	Topics []types.Hash
}

type EventSetTotalFreeFeeLimitSuccess struct {
	Phase  types.Phase
	Amount types.U128
	Topics []types.Hash
}

type EventUnlockStartedFrom struct {
	Phase     types.Phase
	StartFrom types.U32
	Topics    []types.Hash
}

type EventUnlockSuccess struct {
	Phase    types.Phase
	Who      types.AccountID
	UnlockAt types.U32
	Topics   []types.Hash
}

type EventSuperiorChanged struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventMinerChanged struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetLimitSuccess struct {
	Phase  types.Phase
	Limit  types.U128
	Topics []types.Hash
}

type EventMintSuccess struct {
	Phase           types.Phase
	EthereumTxHash  types.H256
	EthereumAddress types.H160
	Amount          types.U128
	Topics          []types.Hash
}

// EventClaimsClaimed is emitted when an account claims some CRUs
type EventClaimsClaimed struct {
	Phase           types.Phase
	Who             types.AccountID
	EthereumAddress types.H160
	Amount          types.U128
	Topics          []types.Hash
}

type EventBondEthSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	EthereumAddress types.H160
	Topics          []types.Hash
}

type EventCru18MinerChanged struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventCru18MintSuccess struct {
	Phase           types.Phase
	EthereumAddress types.H160
	Amount          types.U128
	Topics          []types.Hash
}

type EventCru18Claimed struct {
	Phase           types.Phase
	EthereumAddress types.H160
	Who             types.AccountID
	Amount          types.U128
	Topics          []types.Hash
}

type EventCsmMinerChanged struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventSetCsmLimitSuccess struct {
	Phase  types.Phase
	Limit  types.U128
	Topics []types.Hash
}

type EventCsmMintSuccess struct {
	Phase           types.Phase
	EthereumTxHash  types.H256
	EthereumAddress types.H160
	Amount          types.U128
	Topics          []types.Hash
}

// EventClaimsClaimed is emitted when an account claims some CRUs
type EventCsmClaimed struct {
	Phase           types.Phase
	Who             types.AccountID
	EthereumAddress types.H160
	Amount          types.U128
	Topics          []types.Hash
}

type EventAddBenefitFundsSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Type   FundsType
	Topics []types.Hash
}

type EventCutBenefitFundsSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Type   FundsType
	Topics []types.Hash
}

type EventRebondBenefitFundsSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Type   FundsType
	Topics []types.Hash
}

type EventWithdrawBenefitFundsSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

// DispatchClass is a generalized group of dispatch types. This is only distinguishing normal, user-triggered
// transactions (`Normal`) and anything beyond which serves a higher purpose to the system (`Operational`).
type FundsType struct {
	// A swork funds type
	IsSwork bool
	// An market funds type
	IsMarket bool
}

func (d *FundsType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	switch b {
	case 0:
		d.IsSwork = true
	case 1:
		d.IsMarket = true
	default:
		return fmt.Errorf("unknown FundsType enum: %v", b)
	}
	return err
}

//nolint:gocritic,golint
func (d FundsType) Encode(encoder scale.Encoder) error {
	var err error
	if d.IsSwork {
		err = encoder.PushByte(0)
	} else if d.IsMarket {
		err = encoder.PushByte(1)
	}
	return err
}

type EventFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
	Topics       []types.Hash
}

type EventNonFungibleTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventGenericTransfer struct {
	Phase        types.Phase
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
	Topics       []types.Hash
}

type EventRelayerThresholdChanged struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventChainWhitelisted struct {
	Phase   types.Phase
	ChainId types.U8
	Topics  []types.Hash
}

type EventRelayerAdded struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventRelayerRemoved struct {
	Phase   types.Phase
	Relayer types.AccountID
	Topics  []types.Hash
}

type EventVoteFor struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventVoteAgainst struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Voter        types.AccountID
	Topics       []types.Hash
}

type EventProposalApproved struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalRejected struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalSucceeded struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventProposalFailed struct {
	Phase        types.Phase
	SourceId     types.U8
	DepositNonce types.U64
	Topics       []types.Hash
}

type EventCSMLockingBonded struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCSMLockingUnbonded struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCSMLockingWithdrawn struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCSMLockingSetCSMGuaranteePerfSuccess struct {
	Phase        types.Phase
	Who          types.AccountID
	GuaranteeFee types.U32
	Topics       []types.Hash
}

type EventCSMLockingCSMGuaranteeSuccess struct {
	Phase  types.Phase
	Source types.AccountID
	Target types.AccountID
	Topics []types.Hash
}

type EventCSMLockingCancelCSMGuaranteeSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventCandyCandyIssued struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCandyCandyTransferred struct {
	Phase  types.Phase
	Source types.AccountID
	Target types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventCandyCandyBurned struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type EventBridgeTransferFeeUpdated struct {
	Phase    types.Phase
	ChainId  types.U8
	MinFee   types.U128
	FeeScale types.U32
	Topics   []types.Hash
}

type EventQueueWorkReportSuccess struct {
	Phase       types.Phase
	BlockNumber types.U32
	Index       types.U32
	Who         types.AccountID
	Owner       types.AccountID
	Topics      []types.Hash
}

type EventSetSpowerSuperiorSuccess struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type EventUpdateSpowerSuccess struct {
	Phase        types.Phase
	Who          types.AccountID
	BlockNumber  types.U32
	SworkerCount types.U32
	ProcessCount types.U32
	Topics       []types.Hash
}

type EventUpdateReplicasSuccess struct {
	Phase        types.Phase
	Who          types.AccountID
	BlockNumber  types.U32
	FileCount    types.U32
	ProcessCount types.U32
	Topics       []types.Hash
}
