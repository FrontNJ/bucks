package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bucks/testutil/keeper"
	"bucks/x/bucks/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BucksKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
