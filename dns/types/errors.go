package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidHash  = sdkerrors.Register(ModuleName, 2, "invalid hash")
	ErrInvalidName  = sdkerrors.Register(ModuleName, 3, "invalid name")
	ErrInvalidPath  = sdkerrors.Register(ModuleName, 4, "invalid path")
	ErrInvalidOwner = sdkerrors.Register(ModuleName, 5, "invalid owner")
)
