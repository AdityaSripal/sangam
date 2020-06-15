package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DomainEntry
type DomainEntry interface {
	GetOwner() DomainOwner
	GetSubDomains() []DomainEntry
	GetContents() []ContentEntry
	// returns parent domain, nil if top-level domain
	GetParent() DomainEntry

	// returns the string identifier of this domain
	String() string
	// returns the full path string representation
	// up to the global top-level domain
	Path() string

	// Returns the subdomain under this domain as specified by the given string
	GetSubDomain(path string) DomainEntry

	// Returns the content under this domain as specified by path String
	// Content may either be in this domain or in a subdomain
	// Returns latest version by default
	GetContent(path string) ContentEntry

	// Returns content at a given sequence
	GetContentAtSequenct(path string, seq uint64) ContentEntry

	// Methods to add/update/delete subdomains and content
	AddSubDomain(domain DomainEntry) error
	AddPrecommit(name string, precommit []byte) error
	AddContent(c ContentEntry, reveal uint64) error
	UpdateSubDomain(name string, domain DomainEntry) error
	UpdateContent(newHash []byte) error
	DeleteSubDomain(name string) error
	DeleteContent(name string)
}

// DomainOwner
type DomainOwner interface {
	// authenticates add/update/deletes to subdomains of this domain
	AuthenticateDomainChanges(sdk.Msg) error
	// authenticates add/update/deletes to direct content in this domain
	AuthenticateContentChanges(sdk.Msg) error
}

// ContentEntry
type ContentEntry interface {
	GetName() string
	Path() string
	GetParent() DomainEntry
	GetContentHashes() [][]byte

	// returns content at latest Version
	GetContent() []byte
	GetContentAtVersion(seq uint64) []byte
}
