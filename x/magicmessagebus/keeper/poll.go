package keeper

import (
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreatePoll(ctx sdk.Context, poll types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	b := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(types.KeyPrefix(types.PollKey+poll.Id), b)
}

func (k Keeper) UpdatePoll(ctx sdk.Context, poll types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	b := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(types.KeyPrefix(types.PollKey+poll.Id), b)
}

func (k Keeper) GetPoll(ctx sdk.Context, key string) types.Poll {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	var poll types.Poll
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PollKey+key)), &poll)
	return poll
}

func (k Keeper) HasPoll(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	return store.Has(types.KeyPrefix(types.PollKey + id))
}

func (k Keeper) GetPollOwner(ctx sdk.Context, key string) sdk.AccAddress {
	return k.GetPoll(ctx, key).Creator
}

// DeletePoll deletes a poll
func (k Keeper) DeletePoll(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	store.Delete(types.KeyPrefix(types.PollKey + key))
}

func (k Keeper) GetAllPoll(ctx sdk.Context) (msgs []types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PollKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Poll
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
