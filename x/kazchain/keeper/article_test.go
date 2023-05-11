package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "kazchain/testutil/keeper"
	"kazchain/testutil/nullify"
	"kazchain/x/kazchain/keeper"
	"kazchain/x/kazchain/types"
)

func createNArticle(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Article {
	items := make([]types.Article, n)
	for i := range items {
		items[i].Id = keeper.AppendArticle(ctx, items[i])
	}
	return items
}

func TestArticleGet(t *testing.T) {
	keeper, ctx := keepertest.KazchainKeeper(t)
	items := createNArticle(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetArticle(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestArticleRemove(t *testing.T) {
	keeper, ctx := keepertest.KazchainKeeper(t)
	items := createNArticle(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveArticle(ctx, item.Id)
		_, found := keeper.GetArticle(ctx, item.Id)
		require.False(t, found)
	}
}

func TestArticleGetAll(t *testing.T) {
	keeper, ctx := keepertest.KazchainKeeper(t)
	items := createNArticle(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllArticle(ctx)),
	)
}

func TestArticleCount(t *testing.T) {
	keeper, ctx := keepertest.KazchainKeeper(t)
	items := createNArticle(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetArticleCount(ctx))
}
