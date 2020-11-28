package keeper

import (
	"github.com/allinbits/magic-message-bus/x/magicmessagebus/types"
)

var _ types.QueryServer = Keeper{}
