package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kewr/testutil/keeper"
	"kewr/x/kewr/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.KewrKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
