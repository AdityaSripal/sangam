package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Entry is a value in the DNS with owners and content.
type Entry struct {
	// Owners represent the addresses which have authority over this entry.
	Owners [][]byte `json:"owners" yaml:"owners"`

	// Content represents the data of the entry.
	Content []byte `json:"content" yaml:"content"`
}

// GetOwners returns the owners of the entry.
func (e Entry) GetOwners() [][]byte {
	return e.Owners
}

// GetContent returns the content of the entry.
func (e Entry) GetContent() []byte {
	return e.Content
}

func (e Entry) ValidateBasic() error {
	if len(e.Owners) == 0 {
		return sdkerrors.Wrap("owners cannot be empty")
	}

	if len(e.Content) == 0 {
		return sdkerrors.Wrap("content cannot be empty")
	}

	return nil
}
