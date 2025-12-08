package keeper

import (
	"context"

	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInitResponse{}, nil
}
