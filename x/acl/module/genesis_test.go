package acl_test

import (
	"testing"

	keepertest "github.com/GGEZLabs/vvtxchain/testutil/keeper"
	"github.com/GGEZLabs/vvtxchain/testutil/nullify"
	acl "github.com/GGEZLabs/vvtxchain/x/acl/module"
	"github.com/GGEZLabs/vvtxchain/x/acl/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SuperAdmin: &types.SuperAdmin{
			Address: "41",
		},
		AclAdminList: []types.AclAdmin{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		AclAuthorityList: []types.AclAuthority{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AclKeeper(t)
	acl.InitGenesis(ctx, k, genesisState)
	got := acl.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SuperAdmin, got.SuperAdmin)
	require.ElementsMatch(t, genesisState.AclAdminList, got.AclAdminList)
	require.ElementsMatch(t, genesisState.AclAuthorityList, got.AclAuthorityList)
	// this line is used by starport scaffolding # genesis/test/assert
}
