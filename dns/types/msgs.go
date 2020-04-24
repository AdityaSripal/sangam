package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgPreCommitEntry is the initial registration of a DNS Entry using a commit
// reveal scheme.
type MsgPreCommitEntry struct {
	// Key is the lookup key for the DNS Entry.
	Key []byte `json:"key" yaml:"key"`

	// Hash is the hash of a random nonce and the encoded entry.
	Hash []byte `json:"hash" yaml:"hash"`
}

// GetKey returns the key used for the DNS entry.
func (msg MsgPreCommitEntry) GetKey() []byte {
	return msg.Key
}

// GetHash returns the hash of the nonce and encoded entry.
func (msg MsgPreCommitEntry) GetHash() []byte {
	return msg.Hash
}

// ValidateBasic ensures that neither the key or hash is empty.
func (msg MsgPreCommitEntry) ValidateBasic() error {
	if len(msg.Key) == 0 {
		return sdkerrors.Wrap("key cannot be empty")
	}

	if len(msg.Hash) == 0 {
		return sdkerros.Wrap("hash cannot be empty")
	}

	return nil
}

// MsgCommitEntry is the final registration of a DNS Entry with the
// revealed random nonce and the unhashed Entry.
type MsgCommitEntry struct {
	// Nonce is the random nonce used in the pre-commit message.
	Nonce uint64 `json:"nonce" yaml:"nonce"`

	// Key is the lookup key for the DNS Entry.
	Key []byte `json:"key" yaml:"key"`

	// Entry is the DNS entry being registered.
	Entry Entry `json:"entry" yaml:"entry"`
}

// GetNonce returns the random nonce used in the pre-commit message.
func (msg MsgCommitEntry) GetNonce() uint64 {
	return msg.Nonce
}

// GetKey returns the key used for the DNS entry.
func (msg MsgCommitEntry) GetKey() []byte {
	return msg.Key
}

// GetEntry returns the entry being registered.
func (msg MsgCommitEntry) GetEntry() Entry {
	return msg.Entry
}

// ValidateBasic ensures that the key is not empty and the entry
// is well formed.
func (msg MsgCommitEntry) ValidateBasic() error {
	if len(msg.Key) == 0 {
		return sdkerrors.Wrap("key cannot be 0")
	}

	return msg.Entry.ValidateBasic()
}
