package paramUpgrade

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

// MigrateStore performs in-place store migrations from v3 to v4
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, paramsSubspace *paramstypes.Subspace) error {
	ctx.Logger().Info("MIGRATING STORAGE STORE!")
	// Set the module params

	params := types.NewParams()

	params.ProofWindow = 50

	params.MissesToBurn = 3

	params.MaxContractAgeInBlocks = 100

	params.ChunkSize = 10240

	params.PriceFeed = "jklprice"

	params.PricePerTbPerMonth = 8

	params.DepositAccount = "jkl1arsaayyj5tash86mwqudmcs2fd5jt5zgc3sexc"

	paramsSubspace.SetParamSet(ctx, &params)
	ctx.Logger().Info("DONE MIGRATING!")

	return nil
}
