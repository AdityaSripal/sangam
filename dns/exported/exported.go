package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DomainEntry
type DomainEntry interface {
	GetOwner() DomainOwner
	GetSubDomains() []string
	GetContents() []ContentEntry
	// returns parent domain, nil if top-level domain
	GetParent() string

	// returns the string identifier of this domain
	String() string
	// returns the full path string representation
	// up to the global top-level domain
	Path() string
}

// DomainOwner
type DomainOwner interface {
	// authenticates add/update/deletes to subdomains of this domain
	AuthenticateDomainChanges(sdk.Context, sdk.Msg) error
	// authenticates add/update/deletes to direct content in this domain
	AuthenticateContentChanges(sdk.Context, sdk.Msg) error
}

// ContentEntry
type ContentEntry interface {
	GetName() string
	Path() string
	GetParent() string
	GetContentHashes() [][]byte

	// returns content at latest Version
	GetContent() []byte
	GetContentAtVersion(seq uint64) []byte
}