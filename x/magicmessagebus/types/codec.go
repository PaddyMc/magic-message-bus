package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(MsgCreatePoll{}, "magicmessagebus/CreatePoll", nil)
cdc.RegisterConcrete(MsgUpdatePoll{}, "magicmessagebus/UpdatePoll", nil)
cdc.RegisterConcrete(MsgDeletePoll{}, "magicmessagebus/DeletePoll", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreatePoll{},
	&MsgUpdatePoll{},
	&MsgDeletePoll{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
