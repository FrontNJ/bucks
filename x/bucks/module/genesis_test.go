package bucks_test

import (
	"testing"

	keepertest "bucks/testutil/keeper"
	"bucks/testutil/nullify"
	bucks "bucks/x/bucks/module"
	"bucks/x/bucks/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BucksKeeper(t)
	bucks.InitGenesis(ctx, k, genesisState)
	got := bucks.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
