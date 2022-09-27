package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackal-dao/canine/x/lp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) validateWithdrawLPool(ctx sdk.Context, msg *types.MsgWithdrawLPool) error {
	pool, found := k.GetLPool(ctx, msg.PoolName)

	if !found {
		return types.ErrLiquidityPoolNotFound
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	recordKey := types.LProviderRecordKey(pool.Name, creator.String())
	_, found = k.GetLProviderRecord(ctx, recordKey)

	if !found {
		return types.ErrLProviderRecordNotFound
	}

	if msg.Shares < 0 {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Burn amount cannot be negative",
		)
	}

	return nil
}

func (k msgServer) WithdrawLPool(goCtx context.Context, msg *types.MsgWithdrawLPool) (*types.MsgWithdrawLPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creatorAcc, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.validateWithdrawLPool(ctx, msg); err != nil {
		return nil, err
	}

	pool, _ := k.GetLPool(ctx, msg.PoolName)
	poolCoins := sdk.NewCoins(pool.Coins...)
	totalShares, _ := sdk.NewIntFromString(pool.LPTokenBalance)

	// Calculate tokens to return
	// If LPToken is still locked, take panelty.
	recordKey := types.LProviderRecordKey(pool.Name, creatorAcc.String())
	record, _ := k.GetLProviderRecord(ctx, recordKey)
	
	unlockTime, _ := StringToTime(record.UnlockTime) 

	// This is used to calculate amount of coins to return
	penaltyShare := sdk.NewInt(msg.Shares)

	pm, err := sdk.NewDecFromStr(pool.PenaltyMulti)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to convert penalty" +
		" multiplier; saved in invalid format: %s err: %s",
			pool.PenaltyMulti, err))
	}
	if ctx.BlockTime().Before(unlockTime) {
		penalty := pm.MulInt64(msg.Shares)
		penaltyShare = penaltyShare.Sub(penalty.RoundInt())
	}

	returns, err := CalculatePoolShareBurnReturn(pool, penaltyShare)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"Failed to calculate pool share burn return",	
		)
	}

	burninLToken := sdk.NewCoins(sdk.NewCoin(pool.LptokenDenom, sdk.NewInt(msg.Shares)))

	// Transfer LPToken to module
	sdkErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAcc, types.ModuleName, burninLToken)
	
	if sdkErr != nil {
		return nil, sdkErr
	}

	// Send return coins
	sdkErr = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAcc, returns)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Finally, burn that LPToken :fire:
	sdkErr = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burninLToken)

	if sdkErr != nil {
		return nil, sdkErr
	}

	// Update amount on pool
	poolCoins = poolCoins.Sub(returns)

	totalShares = totalShares.Sub(sdk.NewInt(msg.Shares))

	pool.Coins = poolCoins
	pool.LPTokenBalance= totalShares.String()

	k.SetLPool(ctx, pool)

	return nil, nil
}
