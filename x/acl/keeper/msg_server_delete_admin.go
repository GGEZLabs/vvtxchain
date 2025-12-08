package keeper

import (
	"context"

	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteAdmin(goCtx context.Context, msg *types.MsgDeleteAdmin) (*types.MsgDeleteAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteAdminResponse{}, nil
}
