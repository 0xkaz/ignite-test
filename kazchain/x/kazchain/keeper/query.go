package keeper

import (
	"kazchain/x/kazchain/types"
)

var _ types.QueryServer = Keeper{}
