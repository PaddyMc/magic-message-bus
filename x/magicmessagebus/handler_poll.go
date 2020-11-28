package magicmessagebus

import (
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/keeper"
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/google/uuid"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		Id:      uuid.New().String(),
		Title:   msg.Title,
		Options: msg.Options,
	}

	k.CreatePoll(ctx, poll)

	addr2, _ := sdk.AccAddressFromBech32("cosmos1xm82mkw2jkwkdgq3r0cu8f92r9t2emm8m8xpuw")
	//addr2, _ := sdk.AccAddressFromHex("cosmos1xm82mkw2jkwkdgq3r0cu8f92r9t2emm8m8xpuu")
	sendMsg := banktypes.NewMsgSend(
		msg.Creator,
		addr2,
		sdk.NewCoins(sdk.NewInt64Coin("token", int64(10))),
	)
	k.MessageBus.Publish(banktypes.ModuleName, sendMsg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
