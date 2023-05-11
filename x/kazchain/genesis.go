package kazchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kazchain/x/kazchain/keeper"
	"kazchain/x/kazchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the article
	for _, elem := range genState.ArticleList {
		k.SetArticle(ctx, elem)
	}

	// Set article count
	k.SetArticleCount(ctx, genState.ArticleCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ArticleList = k.GetAllArticle(ctx)
	genesis.ArticleCount = k.GetArticleCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
