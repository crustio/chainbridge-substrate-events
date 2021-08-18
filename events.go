package events

import (
	"fmt"

	"github.com/MyronFanQiu/go-substrate-rpc-client/v3/types"
	"github.com/MyronFanQiu/go-substrate-rpc-client/v3/scale"
)

type Events struct {
	ChainBridge_FungibleTransfer        []EventFungibleTransfer        //nolint:stylecheck,golint
	ChainBridge_NonFungibleTransfer     []EventNonFungibleTransfer     //nolint:stylecheck,golint
	ChainBridge_GenericTransfer         []EventGenericTransfer         //nolint:stylecheck,golint
	ChainBridge_RelayerThresholdChanged []EventRelayerThresholdChanged //nolint:stylecheck,golint
	ChainBridge_ChainWhitelisted        []EventChainWhitelisted        //nolint:stylecheck,golint
	ChainBridge_RelayerAdded            []EventRelayerAdded            //nolint:stylecheck,golint
	ChainBridge_RelayerRemoved          []EventRelayerRemoved          //nolint:stylecheck,golint
	ChainBridge_VoteFor                 []EventVoteFor                 //nolint:stylecheck,golint
	ChainBridge_VoteAgainst             []EventVoteAgainst             //nolint:stylecheck,golint
	ChainBridge_ProposalApproved        []EventProposalApproved        //nolint:stylecheck,golint
	ChainBridge_ProposalRejected        []EventProposalRejected        //nolint:stylecheck,golint
	ChainBridge_ProposalSucceeded       []EventProposalSucceeded       //nolint:stylecheck,golint
	ChainBridge_ProposalFailed          []EventProposalFailed          //nolint:stylecheck,golint
}

type CrustEvents struct {
	Staking_EraReward                       []EventEraReward                    //nolint:stylecheck,golint
	Staking_ValidateSuccess                 []EventValidateSuccess              //nolint:stylecheck,golint
	Staking_GuaranteeSuccess                []EventGuaranteeSuccess             //nolint:stylecheck,golint
	Staking_CutGuaranteeSuccess             []EventCutGuaranteeSuccess          //nolint:stylecheck,golint
	Staking_ChillSuccess                    []EventChillSuccess                 //nolint:stylecheck,golint
	Staking_UpdateStakeLimitSuccess         []EventUpdateStakeLimitSuccess      //nolint:stylecheck,golint
	Swork_RegisterSuccess                   []EventRegisterSuccess              //nolint:stylecheck,golint
	Swork_WorksReportSuccess                []EventWorksReportSuccess           //nolint:stylecheck,golint
	Swork_ABUpgradeSuccess                  []EventABUpgradeSuccess             //nolint:stylecheck,golint
	Swork_SetCodeSuccess                    []EventSetCodeSuccess               //nolint:stylecheck,golint
	Swork_JoinGroupSuccess                  []EventJoinGroupSuccess             //nolint:stylecheck,golint
	Swork_QuitGroupSuccess                  []EventQuitGroupSuccess             //nolint:stylecheck,golint
	Swork_CreateGroupSuccess                []EventCreateGroupSuccess           //nolint:stylecheck,golint
	Swork_KickOutSuccess                    []EventKickOutSuccess               //nolint:stylecheck,golint
	Swork_AddIntoAllowlistSuccess           []EventAddIntoAllowlistSuccess      //nolint:stylecheck,golint
	Swork_RemoveFromAllowlistSuccess        []EventRemoveFromAllowlistSuccess   //nolint:stylecheck,golint
	Swork_SetPunishmentSuccess              []EventSetPunishmentSuccess         //nolint:stylecheck,golint
	Swork_RemoveCodeSuccess                 []EventRemoveCodeSuccess            //nolint:stylecheck,golint
	Market_FileSuccess                      []EventFileSuccess                  //nolint:stylecheck,golint
	Market_RenewFileSuccess                 []EventRenewFileSuccess             //nolint:stylecheck,golint
	Market_AddPrepaidSuccess                []EventAddPrepaidSuccess            //nolint:stylecheck,golint
	Market_CalculateSuccess                 []EventCalculateSuccess             //nolint:stylecheck,golint
	Market_IllegalFileClosed                []EventIllegalFileClosed            //nolint:stylecheck,golint
	Market_RewardMerchantSuccess            []EventRewardMerchantSuccess        //nolint:stylecheck,golint
	Market_SetEnableMarketSuccess           []EventSetEnableMarketSuccess       //nolint:stylecheck,golint
	Market_SetBaseFeeSuccess                []EventSetBaseFeeSuccess            //nolint:stylecheck,golint
	Locks_UnlockStartedFrom                 []EventUnlockStartedFrom            //nolint:stylecheck,golint
	Locks_UnlockSuccess                     []EventUnlockSuccess                //nolint:stylecheck,golint
	Claims_SuperiorChanged                  []EventSuperiorChanged              //nolint:stylecheck,golint
	Claims_MinerChanged                     []EventMinerChanged                 //nolint:stylecheck,golint
	Claims_SetLimitSuccess                  []EventSetLimitSuccess              //nolint:stylecheck,golint
	Claims_MintSuccess                      []EventMintSuccess                  //nolint:stylecheck,golint
	Claims_Claimed                          []EventClaimsClaimed                //nolint:stylecheck,golint
	Benefits_AddBenefitFundsSuccess         []EventAddBenefitFundsSuccess       //nolint:stylecheck,golint
	Benefits_CutBenefitFundsSuccess         []EventCutBenefitFundsSuccess       //nolint:stylecheck,golint
	Benefits_RebondBenefitFundsSuccess      []EventRebondBenefitFundsSuccess    //nolint:stylecheck,golint
	Benefits_WithdrawBenefitFundsSuccess    []EventWithdrawBenefitFundsSuccess  //nolint:stylecheck,golint
}

type EventEraReward struct {
	Phase           types.Phase
	EraIndex        types.U32
	AuthoringReward types.U128
	StakingReward   types.U128
	Topics          []types.Hash
}

type EventValidateSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	GuaranteeFee    types.U32
	Topics          []types.Hash
}

type EventGuaranteeSuccess struct {
	Phase           types.Phase
	Source          types.AccountID
	Target          types.AccountID
	Amount          types.U128
	Topics          []types.Hash
}


type EventCutGuaranteeSuccess struct {
	Phase           types.Phase
	Source          types.AccountID
	Target          types.AccountID
	Amount          types.U128
	Topics          []types.Hash
}


type EventChillSuccess struct {
	Phase           types.Phase
	Stash           types.AccountID
	Controller      types.AccountID
	Topics          []types.Hash
}

type EventUpdateStakeLimitSuccess struct {
	Phase           types.Phase
	Number          types.U32
	Topics          []types.Hash
}

type EventRegisterSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	PubKey          types.Bytes
	Topics          []types.Hash
}

type EventWorksReportSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	PubKey          types.Bytes
	Topics          []types.Hash
}
type EventABUpgradeSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	OldPubKey       types.Bytes
	NewPubKey       types.Bytes
	Topics          []types.Hash
}

type EventSetCodeSuccess struct {
	Phase           types.Phase
	Code            types.Bytes
	BlockNumber     types.U32
	Topics          []types.Hash
}

type EventJoinGroupSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Group           types.AccountID
	Topics          []types.Hash
}

type EventQuitGroupSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Group           types.AccountID
	Topics          []types.Hash
}

type EventCreateGroupSuccess struct {
	Phase           types.Phase
	Group           types.AccountID
	Topics          []types.Hash
}

type EventKickOutSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Topics          []types.Hash
}

type EventAddIntoAllowlistSuccess struct {
	Phase           types.Phase
	Group           types.AccountID
	Who             types.AccountID
	Topics          []types.Hash
}

type EventRemoveFromAllowlistSuccess struct {
	Phase           types.Phase
	Group           types.AccountID
	Who             types.AccountID
	Topics          []types.Hash
}

type EventSetPunishmentSuccess struct {
	Phase           types.Phase
	IsPunishment    bool
	Topics          []types.Hash
}

type EventRemoveCodeSuccess struct {
	Phase           types.Phase
	Code            types.Bytes
	Topics          []types.Hash
}

type EventFileSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Cid             types.Bytes
	Topics          []types.Hash
}

type EventRenewFileSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Cid             types.Bytes
	Topics          []types.Hash
}

type EventAddPrepaidSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Cid             types.Bytes
	Amount          types.U128
	Topics          []types.Hash
}

type EventCalculateSuccess struct {
	Phase           types.Phase
	Cid             types.Bytes
	Topics          []types.Hash
}

type EventIllegalFileClosed struct {
	Phase           types.Phase
	Cid             types.Bytes
	Topics          []types.Hash
}

type EventRewardMerchantSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	Topics          []types.Hash
}

type EventSetEnableMarketSuccess struct {
	Phase           types.Phase
	IsEnable        bool
	Topics          []types.Hash
}

type EventSetBaseFeeSuccess struct {
	Phase           types.Phase
	Amount          types.U128
	Topics          []types.Hash
}

type EventUnlockStartedFrom struct {
	Phase           types.Phase
	StartFrom       types.U32
	Topics          []types.Hash
}

type EventUnlockSuccess struct {
	Phase           types.Phase
	Who             types.AccountID
	UnlockAt        types.U32
	Topics          []types.Hash
}

type EventSuperiorChanged struct {
	Phase           types.Phase
	Who             types.AccountID
	Topics          []types.Hash
}

type EventMinerChanged struct {
	Phase           types.Phase
	Who             types.AccountID
	Topics          []types.Hash
}

type EventSetLimitSuccess struct {
	Phase           types.Phase
	Limit           types.U128
	Topics          []types.Hash
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
	Topics          []types.Hash
}

type EventAddBenefitFundsSuccess struct {
	Phase     types.Phase
	Who       types.AccountID
	Amount    types.U128
	Type      FundsType
	Topics    []types.Hash
}

type EventCutBenefitFundsSuccess struct {
	Phase     types.Phase
	Who       types.AccountID
	Amount    types.U128
	Type      FundsType
	Topics    []types.Hash
}

type EventRebondBenefitFundsSuccess struct {
	Phase     types.Phase
	Who       types.AccountID
	Amount    types.U128
	Type      FundsType
	Topics    []types.Hash
}

type EventWithdrawBenefitFundsSuccess struct {
	Phase     types.Phase
	Who       types.AccountID
	Amount    types.U128
	Topics    []types.Hash
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
