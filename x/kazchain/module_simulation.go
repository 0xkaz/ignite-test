package kazchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"kazchain/testutil/sample"
	kazchainsimulation "kazchain/x/kazchain/simulation"
	"kazchain/x/kazchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = kazchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateArticle = "op_weight_msg_article"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateArticle int = 100

	opWeightMsgUpdateArticle = "op_weight_msg_article"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateArticle int = 100

	opWeightMsgDeleteArticle = "op_weight_msg_article"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteArticle int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kazchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ArticleList: []types.Article{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ArticleCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kazchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateArticle int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateArticle, &weightMsgCreateArticle, nil,
		func(_ *rand.Rand) {
			weightMsgCreateArticle = defaultWeightMsgCreateArticle
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateArticle,
		kazchainsimulation.SimulateMsgCreateArticle(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateArticle int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateArticle, &weightMsgUpdateArticle, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateArticle = defaultWeightMsgUpdateArticle
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateArticle,
		kazchainsimulation.SimulateMsgUpdateArticle(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteArticle int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteArticle, &weightMsgDeleteArticle, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteArticle = defaultWeightMsgDeleteArticle
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteArticle,
		kazchainsimulation.SimulateMsgDeleteArticle(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
