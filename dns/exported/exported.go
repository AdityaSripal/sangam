package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Entry is the entry used in the DNS.
type Entry interface {
	GetOwners() []sdk.AccAddress
	GetContent() []byte
	GetSequence() uint64
}

// Domain is the lookup key used in the DNS.
type Domain interface {
	GetPrefix() string
	GetContentName() string
	GetBytes() []byte
}
