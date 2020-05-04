package dns

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns sdk.Handler for DNS module messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgPreCommitEntry:
			return handleMsgPreCommitEntry(ctx, k, msg)
		case MsgCommitEntry:
			return handleMsgCommitEntry(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ICS-20 transfer message type: %T", msg)
		}
	}
}

// handleMsgPreCommitEntry defines the sdk.Handler for MsgPreCommitEntry.
func handleMsgPreCommitEntry(ctx sdk.Context, k Keeper, msg MsgPreCommitEntry) (*sdk.Result, error) {
	err := k.StorePreCommit(ctx, msg.GetPath(), msg.GetHash())
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// handleMsgCommitEntry defines the sdk.Handler for MsgCommitEntry.
func handleMsgCommitEntry(ctx sdk.Context, k Keeper, msg MsgCommitEntry) (*sdk.Result, error) {
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
