package keeper

import (
	"github.com/GGEZLabs/vvtxchain/x/trade/types"
)

var _ types.QueryServer = Keeper{}
