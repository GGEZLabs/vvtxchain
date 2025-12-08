package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/GGEZLabs/vvtxchain/testutil/keeper"
	"github.com/GGEZLabs/vvtxchain/x/trade/keeper"
	"github.com/GGEZLabs/vvtxchain/x/trade/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(tb testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	tb.Helper()
	k, ctx := keepertest.TradeKeeper(tb)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
