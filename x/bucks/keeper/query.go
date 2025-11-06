package keeper

import (
	"bucks/x/bucks/types"
)

var _ types.QueryServer = Keeper{}
