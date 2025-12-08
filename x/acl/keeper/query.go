package keeper

import (
	"github.com/GGEZLabs/vvtxchain/x/acl/types"
)

var _ types.QueryServer = Keeper{}
