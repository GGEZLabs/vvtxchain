package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/GGEZLabs/vvtxchain/testutil/keeper"
	"github.com/GGEZLabs/vvtxchain/x/acl/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AclKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
