package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Default sequence values for a pre-commit and commit entry
const (
	PreCommitSequence = 0
	CommitSequence    = 1
)

// Entry is a value in the DNS with owners and a hash of the content being
// registered.
type Entry struct {
	// Owners represent the addresses which have authority over this entry.
	Owners []sdk.AccAddress `json:"owners" yaml:"owners"`

	// ContentHash represents the hash of the entry's data.
	ContentHash []byte `json:"content_hash" yaml:"content_hash"`

	// Sequence
	Sequence uint64 `json:"sequence" yaml:"sequence"`
}

// GetOwners returns the owners of the entry.
func (e Entry) GetOwners() []sdk.AccAddress {
	return e.Owners
}

// GetContentHash returns the hash of the content.
func (e Entry) GetContentHash() []byte {
	return e.ContentHash
}

// GetSequence returns the current sequence number.
func (e Entry) GetSequence() uint64 {
	return e.Sequence
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
		return sdkerrors.Wrap(ErrInvalidOwners, "owners cannot be empty")
	}

	if len(e.ContentHash) == 0 {
		return sdkerrors.Wrap(ErrInvalidContentHash, "content hash cannot be empty")
	}

	return nil
}
