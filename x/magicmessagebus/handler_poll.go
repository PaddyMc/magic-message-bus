package magicmessagebus

import (
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/keeper"
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		Id:      uuid.New().String(),
		Title:   msg.Title,
	}

	k.CreatePoll(ctx, poll)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
