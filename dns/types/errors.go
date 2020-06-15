package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidName       = sdkerrors.Register(ModuleName, 2, "invalid name")
	ErrEntryDoesNotExist = sdkerrors.Register(ModuleName, 3, "entry does not exist")
	ErrEntryExists       = sdkerrors.Register(ModuleName, 4, "entry already exists")
)
