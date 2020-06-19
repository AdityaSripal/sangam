package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AdityaSripal/sangam/dns/exported"
)

// message types for DNS module
const (
	TypeMsgRegisterDomain string = "register_domain"
	TypeMsgPreCommitEntry string = "pre_commit_entry"
	TypeMsgCommitEntry    string = "commit_entry"
)

// ensure messages implement sdk.Msg
var (
	_ sdk.Msg = MsgRegisterDomain{}
	_ sdk.Msg = MsgPreCommitEntry{}
	_ sdk.Msg = MsgCommitEntry{}
)

// MsgRegisterDomain is used to register a new domain that does not already exist. This message
// does not require a commit-reveal scheme
type MsgRegisterDomain struct {
	Name       string
	ParentPath string
	Owner      exported.DomainOwner
	Signers    []sdk.AccAddress
}

// NewMsgRegisterDomain returns a new instance of MsgRegisterDomain.
func NewMsgRegisterDomain(name, parentPath string, owner exported.DomainOwner, signers []sdk.AccAddress) MsgRegisterDomain {
	return MsgRegisterDomain{
		Name:       name,
		ParentPath: parentPath,
		Owner:      owner,
		Signers:    signers,
	}
}

// Domain returns the domain name for this domain.
func (msg MsgRegisterDomain) GetDomain() string {
	return msg.Name
}

// GetParentPath returns the full path of the parents. An empty string is used if the domain
// being registered is a top-level domain.
func (msg MsgRegisterDomain) GetParentPath() string {
	return msg.ParentPath
}

// GetOwner returns the owner object of this domain.
func (msg MsgRegisterDomain) GetOwner() exported.DomainOwner {
	return msg.Owner
}

// Route implements sdk.Msg
func (msg MsgRegisterDomain) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgRegisterDomain) Type() string {
	return TypeMsgRegisterDomain
}

// ValidateBasic does basic validation of the message fields.
func (msg MsgRegisterDomain) ValidateBasic() error {
	if strings.TrimSpace(msg.Name) == "" {
		return sdkerrors.Wrap(ErrInvalidName, "domain name cannot be empty")
	}
	if strings.TrimSpace(msg.ParentPath) == "" {
		return sdkerrors.Wrap(ErrInvalidPath, "domain parent path cannot be empty")
	}
	if msg.Owner == nil {
		return sdkerrors.Wrap(ErrInvalidOwner, "domain owner cannot be empty")
	}
	if len(msg.Signers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "there must be at least one signer")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgRegisterDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterDomain) GetSigners() []sdk.AccAddress {
	return msg.Signers
}

// MsgPreCommitEntry is the initial registration of a DNS Entry using a commit
// reveal scheme.
// Hash = hash(content_hash + domain_path + random_nonce)
type MsgPreCommitEntry struct {
	Path        string           `json:"domain_path" yaml:"domain_path"`
	ContentName string           `json:"content_name" yaml:"content_name"`
	Hash        []byte           `json:"hash" yaml:"hash"`
	Signers     []sdk.AccAddress `json:"signers" yaml:"signers"`
}

// NewMsgPreCommitEntry returns a new instance of MsgPreCommitEntry.
func NewMsgPreCommitEntry(domainPath, contentName string, hash []byte, signers []sdk.AccAddress) MsgPreCommitEntry {
	return MsgPreCommitEntry{
		Path:        domainPath,
		ContentName: contentName,
		Hash:        hash,
		Signers:     signers,
	}
}

// DomainPath returns the fill path of the domain.
func (msg MsgPreCommitEntry) DomainPath() string {
	return msg.Path
}

// Name returns the name of the content being registered.
func (msg MsgPreCommitEntry) Name() string {
	return msg.ContentName
}

// GetHash returns the hash of the random nonce + encoded entry
func (msg MsgPreCommitEntry) GetHash() []byte {
	return msg.Hash
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
	if len(msg.Signers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Signers cannot be empty")
	}

	return nil
}

// GetSignBytes implements Msg
func (msg MsgPreCommitEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgPreCommitEntry) GetSigners() []sdk.AccAddress {
	return msg.Signers
}

// MsgCommitEntry is the final registration of a DNS Entry with the
// revealed random nonce and the unhashed Entry.
type MsgCommitEntry struct {
	Nonce       uint64           `json:"nonce" yaml:"nonce"`               // random nonce used in pre-commit message
	ContentHash []byte           `json:"content_hash" yaml:"content_hash"` // hash of the content being registered
	DomainPath  string           `json:"domain_path" yaml:"domain_path"`   // full path of the domain
	Name        string           `json:"name" yaml:"name"`                 // name of the content
	Signers     []sdk.AccAddress `json:"signers" yaml:"signers"`           // signers of this message
}

// NewMsgCommitEntry returns a new instance of MsgCommitEntry.
func NewMsgCommitEntry(nonce uint64, contentHash []byte, domainPath, name string, signers []sdk.AccAddress) MsgCommitEntry {
	return MsgCommitEntry{
		Nonce:       nonce,
		ContentHash: contentHash,
		DomainPath:  domainPath,
		Name:        name,
		Signers:     signers,
	}
}

// Nonce returns the random nonce used in the pre-commit message.
func (msg MsgCommitEntry) GetNonce() uint64 {
	return msg.Nonce
}

// ContentHash returns the content hash being registered.
func (msg MsgCommitEntry) GetContentHash() []byte {
	return msg.ContentHash
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
	if len(msg.ContentHash) == 0 {
		return sdkerrors.Wrap(ErrInvalidHash, "no hash of content")
	}
	if len(msg.Signers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "owners cannot be empty")
	}
	return nil
}

// GetSignBytes implements Msg
func (msg MsgCommitEntry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg MsgCommitEntry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signers[0]}
}
