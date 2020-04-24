package dns

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// HandleMsgPreCommitEntry defines the sdk.Handler for MsgPreCommitEntry
func HandleMsgPreCommitEntry(ctx sdk.Context, k Keeper, msg exported.MsgPreCommitEntry) (*sdk.Result, error) {
	err := k.StorePreCommit(ctx, msg.GetKey(), msg.GetHash())
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
