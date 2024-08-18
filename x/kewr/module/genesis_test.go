package kewr_test

import (
	"testing"

	keepertest "kewr/testutil/keeper"
	"kewr/testutil/nullify"
	kewr "kewr/x/kewr/module"
	"kewr/x/kewr/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KewrKeeper(t)
	kewr.InitGenesis(ctx, k, genesisState)
	got := kewr.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
