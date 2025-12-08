package keeper

import (
	"context"

	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAuthority(goCtx context.Context, msg *types.MsgAddAuthority) (*types.MsgAddAuthorityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddAuthorityResponse{}, nil
}
