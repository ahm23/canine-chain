package keeper

import (
	"github.com/jackalLabs/canine-chain/v4/x/jklmint/types"
)

var _ types.QueryServer = Keeper{}
