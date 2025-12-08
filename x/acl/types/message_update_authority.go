package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAuthority{}

func NewMsgUpdateAuthority(creator string, authAddress string, newName string, overwriteAccessDefinitions string, addAccessDefinitions string, updateAccessDefinition string, deleteAccessDefinitions []string, clearAllAccessDefinitions bool) *MsgUpdateAuthority {
	return &MsgUpdateAuthority{
		Creator:                    creator,
		AuthAddress:                authAddress,
		NewName:                    newName,
		OverwriteAccessDefinitions: overwriteAccessDefinitions,
		AddAccessDefinitions:       addAccessDefinitions,
		UpdateAccessDefinition:     updateAccessDefinition,
		DeleteAccessDefinitions:    deleteAccessDefinitions,
		ClearAllAccessDefinitions:  clearAllAccessDefinitions,
	}
}

func (msg *MsgUpdateAuthority) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
