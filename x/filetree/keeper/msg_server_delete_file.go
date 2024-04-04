package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/filetree/types"
)

func (k msgServer) DeleteFile(goCtx context.Context, msg *types.MsgDeleteFile) (*types.MsgDeleteFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ownerAddress := MakeOwnerAddress(msg.HashPath, msg.Account)

	file, found := k.GetFiles(ctx, msg.HashPath, ownerAddress)
	if !found {
		return nil, types.ErrFileNotFound
	}
	// Only the owner of a file can delete it
	isOwner := IsOwner(file, msg.Creator)
	if !isOwner {
		return nil, types.ErrCannotDelete
	}

	k.RemoveFiles(ctx, msg.HashPath, ownerAddress)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	return &types.MsgDeleteFileResponse{}, nil
}
