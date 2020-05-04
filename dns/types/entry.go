package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Entry is a value in the DNS with owners and content.
type Entry struct {
	// Owners represent the addresses which have authority over this entry.
	Owners []sdk.AccAddress `json:"owners" yaml:"owners"`

	// ContentHash represents the hash of the entry's data.
	ContentHash []byte `json:"content" yaml:"content"`
}

// GetOwners returns the owners of the entry.
func (e Entry) GetOwners() []sdk.AccAddress {
	return e.Owners
}

// GetContentHash returns the hash of the content.
func (e Entry) GetContentHash() []byte {
	return e.Content
}

// GetBytes returns the json encoded entry.
func (e Entry) GetBytes() []byte {
	bz, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic verifies that the owners and content hash are not empty.
func (e Entry) ValidateBasic() error {
	if len(e.Owners) == 0 {
		return sdkerrors.Wrap("owners cannot be empty")
	}

	if len(e.Content) == 0 {
		return sdkerrors.Wrap("content cannot be empty")
	}

	return nil
}
