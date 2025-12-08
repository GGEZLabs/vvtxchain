package keeper

import (
	"context"

	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateSuperAdmin(goCtx context.Context, msg *types.MsgUpdateSuperAdmin) (*types.MsgUpdateSuperAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateSuperAdminResponse{}, nil
}
