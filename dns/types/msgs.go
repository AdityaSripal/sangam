package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Message types for DNS module
const (
	TypeMsgPreCommitEntry string = "pre_commit_entry"
	TypeMsgCommitEntry    string = "commit_entry"
)

var _ sdk.Msg = MsgPreCommitEntry{}

// MsgPreCommitEntry is the initial registration of a DNS Entry using a commit
// reveal scheme.
type MsgPreCommitEntry struct {
	// Domain is the lookup key for the DNS Entry.
	Domain Domain `json:"path" yaml:"path"`

	// Hash is the hash of a random nonce and the encoded entry.
	Hash []byte `json:"hash" yaml:"hash"`

	// owners of the entry, the first account is expected to be the signer
	// of this message.
	Owners []sdk.AccAddress `json:"owners" yaml:"owners"`
}

// NewMsgPreCommitEntry returns a new instance of MsgPreCommitEntry.
func NewMsgPreCommitEntry(domain Domain, hash []byte, owners []sdk.AccAddress) MsgPreCommitEntry {
	return MsgPreCommitEntry{
		Domain: domain,
		Hash:   hash,
		Owners: owners,
	}
}

// Route implements sdk.Msg
func (msg MsgPreCommitEntry) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgPreCommitEntry) Type() string {
	return TypeMsgPreCommitEntry
}

// ValidateBasic ensures that hash is not empty and the path is valid.
func (msg MsgPreCommitEntry) ValidateBasic() error {
	if len(msg.Hash) == 0 {
		return sdkerrors.Wrap(ErrInvalidHash, "hash cannot be empty")
	}
	if len(msg.Owners) == 0 {
		return sdkerrors.Wrap(ErrInvalidOwners, "pre commit entry message must have at least one owner")
	}

	return msg.Domain.ValidateBasic()
}

// GetSignBytes implements Msg
func (msg MsgPreCommitEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgPreCommitEntry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owners[0]}
}

// GetDomain returns the path used to map the DNS entry.
func (msg MsgPreCommitEntry) GetDomain() Domain {
	return msg.Domain
}

// GetOwners returns the owners of the future entry.
func (msg MsgPreCommitEntry) GetOwners() []sdk.AccAddress {
	return msg.Owners
}

// GetHash returns the hash of the random nonce + encoded entry
func (msg MsgPreCommitEntry) GetHash() []byte {
	return msg.Hash
}

// MsgCommitEntry is the final registration of a DNS Entry with the
// revealed random nonce and the unhashed Entry.
type MsgCommitEntry struct {
	// Nonce is the random nonce used in the pre-commit message.
	Nonce uint64 `json:"nonce" yaml:"nonce"`

	// Domain is the lookup key for the DNS Entry.
	Domain Domain `json:"domain" yaml:"domain"`

	// Entry is the DNS entry being registered.
	Entry Entry `json:"entry" yaml:"entry"`

	// Owners of the commit entry. First account is expected to be the signer
	// of this message. Owners must match the pre-commit
	Owners []sdk.AccAddress
}

// NewMsgCommitEntry returns a new instance of MsgCommitEntry.
func NewMsgCommitEntry(nonce uint64, domain Domain, entry Entry, owners []sdk.AccAddress) MsgCommitEntry {
	return MsgCommitEntry{
		Nonce:  nonce,
		Domain: domain,
		Entry:  entry,
		Owners: owners,
	}
}

// Route implements sdk.Msg.
func (msg MsgCommitEntry) Route() string {
	return RouterKey
}

// Type implements sdk.Msg.
func (msg MsgCommitEntry) Type() string {
	return TypeMsgCommitEntry
}

// ValidateBasic ensures that the key is not empty and the entry
// is well formed.
func (msg MsgCommitEntry) ValidateBasic() error {
	if err := msg.Domain.ValidateBasic(); err != nil {
		return err
	}
	if len(msg.Owners) == 0 {
		return sdkerrors.Wrap(ErrInvalidOwners, "owners cannot be empty")
	}

	return msg.Entry.ValidateBasic()
}

// GetSignBytes implements Msg
func (msg MsgCommitEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgCommitEntry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owners[0]}
}

// GetNonce returns the random nonce used in the pre-commit message.
func (msg MsgCommitEntry) GetNonce() uint64 {
	return msg.Nonce
}

// GetDomain returns the key used for the DNS entry.
func (msg MsgCommitEntry) GetDomain() Domain {
	return msg.Domain
}

// GetEntry returns the entry being registered.
func (msg MsgCommitEntry) GetEntry() Entry {
	return msg.Entry
}
