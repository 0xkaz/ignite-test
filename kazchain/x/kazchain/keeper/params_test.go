package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "kazchain/testutil/keeper"
	"kazchain/x/kazchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KazchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
