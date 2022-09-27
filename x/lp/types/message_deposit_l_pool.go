package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDepositLPool = "deposit_l_pool"

var _ sdk.Msg = &MsgDepositLPool{}

func NewMsgDepositLPool(creator string, poolName string, coins sdk.DecCoins, lockDuration int64) *MsgDepositLPool {
	return &MsgDepositLPool{
		Creator:  creator,
		PoolName: poolName,
		Coins:    coins,
		LockDuration: lockDuration, 
	}
}

func (msg *MsgDepositLPool) Route() string {
	return RouterKey
}

func (msg *MsgDepositLPool) Type() string {
	return TypeMsgDepositLPool
}

func (msg *MsgDepositLPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDepositLPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositLPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	coins := sdk.NormalizeCoins(msg.Coins)
	if err != nil {
		return err
	}

	if !coins.IsValid() {
		return sdkerrors.ErrInvalidCoins
	}

	if msg.LockDuration < 0 {
		return errors.New(fmt.Sprintf("Lock duration cannot be less than 0: %v",
			msg.LockDuration))
	}

	return nil
}
