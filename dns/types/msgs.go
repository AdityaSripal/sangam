package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgPreCommitEntry is the initial registration of a DNS Entry using a commit
// reveal scheme.
type MsgPreCommitEntry struct {
	// Path is the lookup key for the DNS Entry.
	Path Path `json:"path" yaml:"path"`

	// Hash is the hash of a random nonce and the encoded entry.
	Hash []byte `json:"hash" yaml:"hash"`
}

// GetPath returns the path used to map the DNS entry.
func (msg MsgPreCommitEntry) GetPath() []byte {
	return msg.Path
}

// GetHash returns the hash of the nonce and encoded entry.
func (msg MsgPreCommitEntry) GetHash() []byte {
	return msg.Hash
}

// ValidateBasic ensures that hash is not empty and the path is valid.
func (msg MsgPreCommitEntry) ValidateBasic() error {
	if len(msg.Hash) == 0 {
		return sdkerros.Wrap("hash cannot be empty")
	}

	return msg.Path.ValidateBasic()
}

// MsgCommitEntry is the final registration of a DNS Entry with the
// revealed random nonce and the unhashed Entry.
type MsgCommitEntry struct {
	// Nonce is the random nonce used in the pre-commit message.
	Nonce uint64 `json:"nonce" yaml:"nonce"`

	// Path is the lookup key for the DNS Entry.
	Path []byte `json:"path" yaml:"path"`

	// Entry is the DNS entry being registered.
	Entry Entry `json:"entry" yaml:"entry"`
}

// GetNonce returns the random nonce used in the pre-commit message.
func (msg MsgCommitEntry) GetNonce() uint64 {
	return msg.Nonce
}

// GetPath returns the key used for the DNS entry.
func (msg MsgCommitEntry) GetPath() []byte {
	return msg.Path
}

// GetEntry returns the entry being registered.
func (msg MsgCommitEntry) GetEntry() Entry {
	return msg.Entry
}

// ValidateBasic ensures that the key is not empty and the entry
// is well formed.
func (msg MsgCommitEntry) ValidateBasic() error {
	if err := msg.Path.ValidateBasic(); err != nil {
		return err
	}

	return msg.Entry.ValidateBasic()
}
