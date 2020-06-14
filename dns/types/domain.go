package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DomainEntry represents the key used to map to a DNS entry
type DomainEntry struct {
	Owner      DomainOwner `json:"owner" yaml:"owner"`
	SubDomains []DomainEntry
	Contents   []ContentEntry
	Parent     DomainEntry

	Name string `json:"name" yaml:"name"`
}

// Owner returns the domain owner.
func (de DomainEntry) Owner() DomainOwner {
	return de.Owner
}

// SubDomains returns the sub-domains underneath this domain entry.
func (de DomainEntry) SubDomains() string {
	return de.SubDomains
}

// Contents returns the contents represented at this domain entry.
func (de DomainEntry) Contents() []ContentryEntry {
	return de.Contents
}

// Parent returns the parent entry for this domain entry.
func (de DomainEntry) Parents() DomainEntry {
	return de.Parent
}

// String returns the string identifier for this domain.
func (de DomainEntry) String() string {
	return Name
}

// Path returns the full path string representation up to the global top-level
func (de DomainEntry) Path() string {
	return de.Parent.Path() + de.String()
}

// GetSubDomain returns the subdomain under this domain as specified by the given string.
func (de DomainEntry) GetSubDomain(path string) DomainEntry {
	// TODO: implement
	return nil
}

// GetContent returns the content under this domain as specified by the path. The content
// may either be in this domain or in a sub-domain. The latest version is returned by default.
func (de DomainEntry) GetContent(path string) ContentEntry {
	// TODO: implement
	return nil
}

// GetContentAtSequence returns the content at the given sequence.
func (de DomainEntry) GetContentAtSequence(path string, seq uint64) ContentEntry {
	// TODO: implement
	return nil
}

// AddSubDomain adds a sub-domain to the current domain entry.
func (de DomainEntry) AddSubDomain(domain DomainEntry) error {
	if de.GetSubDomain(domain) != nil {
		return sdkerrors.Wrapf(
			ErrEntryExists,
			"cannot add sub-domain (%v) that already exists within this domain (%v)", domain, de,
		)
	}

	de.SubDomains = append(de.SubDomains, domain)
	return nil
}

// AddPreCommit adds a pre-commit related to this domain entry.
func (de DomainEntry) AddPreCommit(name string, precommit []byte) error {
	// TODO: implement
	return nil
}

// AddContent adds a content entry to this domain.
func (de DomainEntry) AddContent(content ContentEntry, reveal uint64) error {
	// TODO: implement

	if de.GetContent(content) != nil {
		return sdkerrors.Wrapf(
			ErrEntryExists,
			"cannot add content (%v) that already exists within this domain (%v)", content, de,
		)
	}

	de.Contents = append(de.Contents, content)
	return nil
}

// UpdateSubDomain updates the sub-domain contained within this entry.
