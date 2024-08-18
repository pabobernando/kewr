package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kewr/testutil/keeper"
	"kewr/x/kewr/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.KewrKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
