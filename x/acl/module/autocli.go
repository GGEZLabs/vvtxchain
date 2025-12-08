package acl

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/GGEZLabs/vvtxchain/api/vvtxchain/acl"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SuperAdmin",
					Use:       "show-super-admin",
					Short:     "show super_admin",
				},
				{
					RpcMethod: "AclAdminAll",
					Use:       "list-acl-admin",
					Short:     "List all acl_admin",
				},
				{
					RpcMethod:      "AclAdmin",
					Use:            "show-acl-admin [id]",
					Short:          "Shows a acl_admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "AclAuthorityAll",
					Use:       "list-acl-authority",
					Short:     "List all acl_authority",
				},
				{
					RpcMethod:      "AclAuthority",
					Use:            "show-acl-authority [id]",
					Short:          "Shows a acl_authority",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Init",
					Use:            "init [super-admin]",
					Short:          "Send a init tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "superAdmin"}},
				},
				{
					RpcMethod:      "AddAdmin",
					Use:            "add-admin [admins]",
					Short:          "Send a add_admin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "admins"}},
				},
				{
					RpcMethod:      "DeleteAdmin",
					Use:            "delete-admin [admins]",
					Short:          "Send a delete_admin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "admins"}},
				},
				{
					RpcMethod:      "AddAuthority",
					Use:            "add-authority [auth-address] [name] [access-definitions]",
					Short:          "Send a add_authority tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "authAddress"}, {ProtoField: "name"}, {ProtoField: "accessDefinitions"}},
				},
				{
					RpcMethod:      "DeleteAuthority",
					Use:            "delete-authority [auth-address]",
					Short:          "Send a delete_authority tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "authAddress"}},
				},
				{
					RpcMethod:      "UpdateAuthority",
					Use:            "update-authority [auth-address] [new-name] [overwrite-access-definitions] [add-access-definitions] [update-access-definition] [delete-access-definitions] [clear-all-access-definitions]",
					Short:          "Send a update_authority tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "authAddress"}, {ProtoField: "newName"}, {ProtoField: "overwriteAccessDefinitions"}, {ProtoField: "addAccessDefinitions"}, {ProtoField: "updateAccessDefinition"}, {ProtoField: "deleteAccessDefinitions"}, {ProtoField: "clearAllAccessDefinitions"}},
				},
				{
					RpcMethod:      "UpdateSuperAdmin",
					Use:            "update-super-admin [new-super-admin]",
					Short:          "Send a update_super_admin tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "newSuperAdmin"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
