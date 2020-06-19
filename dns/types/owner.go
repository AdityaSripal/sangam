package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type SimpleOwner struct {
	Address sdk.Address
}

func (so SimpleOwner) AuthenticateDomainChanges(_ sdk.Context, msg sdk.Msg) error {
	for _, signer := range msg.GetSigners() {
		if signer.Equals(so.Address) {
			return nil
		}
	}
	return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "owner did not authenticate domain changes")
}

func (so SimpleOwner) AuthenticateContentChanges(_ sdk.Context, msg sdk.Msg) error {
	for _, signer := range msg.GetSigners() {
		if signer.Equals(so.Address) {
			return nil
		}
	}
	return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "owner did not authenticate domain changes")
}
