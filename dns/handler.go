package dns

import (
	"bytes"
	"crypto/sha256"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	entry := Entry{
		Owners:      msg.GetOwners(),
		ContentHash: msg.GetHash(),
		Sequence:    PreCommitSequence, // 0
	}
	k.SetEntry(ctx, msg.GetDomain(), entry)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}

// handleMsgCommitEntry defines the sdk.Handler for MsgCommitEntry.
func handleMsgCommitEntry(ctx sdk.Context, k Keeper, msg MsgCommitEntry) (*sdk.Result, error) {
	preCommitEntry, found := k.GetEntry(ctx, msg.GetDomain())
	if !found {
		return nil, sdkerrors.Wrapf(ErrEntryNotFound, "pre-commit entry not found for domain %s", msg.GetDomain())
	}

	hasher := sha256.New()
	value := append(sdk.Uint64ToBigEndian(msg.GetNonce()), msg.GetEntry().GetBytes()...)
	if bytes.Equal(preCommitEntry.GetContentHash(), hasher.Sum(value)) {
		return nil, sdkerrors.Wrapf(ErrHashDoesNotMatch, "pre-commit entry hash %v does not match hash for entry %v under the domain %s", preCommitEntry.GetContentHash(), msg.GetEntry(), msg.GetDomain())
	}

	k.SetEntry(ctx, msg.GetDomain(), msg.GetEntry())

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
