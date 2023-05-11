package kazchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "kazchain/testutil/keeper"
	"kazchain/testutil/nullify"
	"kazchain/x/kazchain"
	"kazchain/x/kazchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ArticleList: []types.Article{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ArticleCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KazchainKeeper(t)
	kazchain.InitGenesis(ctx, *k, genesisState)
	got := kazchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ArticleList, got.ArticleList)
	require.Equal(t, genesisState.ArticleCount, got.ArticleCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
