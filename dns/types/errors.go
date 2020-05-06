package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidOwners       = sdkerrors.Register(ModuleName, 2, "invalid owners")
	ErrInvalidContentHash  = sdkerrors.Register(ModuleName, 3, "invalid content hash")
	ErrInvalidHash         = sdkerrors.Register(ModuleName, 4, "invalid hash")
	ErrInvalidContentName  = sdkerrors.Register(ModuleName, 5, "invalid content name")
	ErrInvalidDomainString = sdkerrors.Register(ModuleName, 6, "invalid domain string")
	ErrEntryNotFound       = sdkerrors.Register(ModuleName, 7, "entry not found")
	ErrHashDoesNotMatch    = sdkerrors.Register(ModuleName, 8, "derived hash does not match expected hash")
)
