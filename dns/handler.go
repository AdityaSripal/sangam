package dns

import (
	"github.com/adityasripal/sangam/dns/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// HandleMsgPreCommitEntry defines the sdk.Handler for MsgPreCommitEntry.
func HandleMsgPreCommitEntry(ctx sdk.Context, k Keeper, msg types.MsgPreCommitEntry) (*sdk.Result, error) {
	err := k.StorePreCommit(ctx, msg.GetPath(), msg.GetHash())
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// HandleMsgCommitEntry defines the sdk.Handler for MsgCommitEntry.
func HandleMsgCommitEntry(ctx sdk.Context, k Keeper, msg types.MsgCommitEntry) (*sdk.Result, error) {
	preCommitHash, found := k.GetPreCommit(ctx, msg.GetPath())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrPreCommitNotFound)
	}

	hash := sha256.New()
	hash.Write(sdk.Uint64ToBigEndian(msg.GetNonce()))
	hash.Write(msg.GetEntry.GetBytes())
	if preCommitHash != hash.Sum(nil) {
		return nil, sdkerrors.Wrap(types.ErrHashDoesNotMatch)
	}

	err = k.SetEntry(ctx, msg.GetPath, msg.GetEntry())
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
