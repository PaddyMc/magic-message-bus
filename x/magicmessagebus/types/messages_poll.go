package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgCreatePoll(creator sdk.AccAddress, title string) *MsgCreatePoll {
  return &MsgCreatePoll{
		Creator: creator,
    Title: title,
	}
}

func (msg *MsgCreatePoll) Route() string {
  return RouterKey
}

func (msg *MsgCreatePoll) Type() string {
  return "CreatePoll"
}

func (msg *MsgCreatePoll) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgCreatePoll) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePoll) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgUpdatePoll{}

func NewMsgUpdatePoll(creator sdk.AccAddress, id string, title string) *MsgUpdatePoll {
  return &MsgUpdatePoll{
        Id: id,
		Creator: creator,
    Title: title,
	}
}

func (msg *MsgUpdatePoll) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePoll) Type() string {
  return "UpdatePoll"
}

func (msg *MsgUpdatePoll) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgUpdatePoll) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePoll) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}

var _ sdk.Msg = &MsgCreatePoll{}

func NewMsgDeletePoll(creator sdk.AccAddress, id string) *MsgDeletePoll {
  return &MsgDeletePoll{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeletePoll) Route() string {
  return RouterKey
}

func (msg *MsgDeletePoll) Type() string {
  return "DeletePoll"
}

func (msg *MsgDeletePoll) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgDeletePoll) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePoll) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
