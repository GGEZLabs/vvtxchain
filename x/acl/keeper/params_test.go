package keeper_test

import (
	"testing"

	keepertest "github.com/GGEZLabs/vvtxchain/testutil/keeper"
	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AclKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.Equal(t, params, k.GetParams(ctx))
}
