package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Entry is the entry used in the DNS.
type Entry interface {
	GetOwners() []sdk.AccAddress
	GetContent() []byte
}

// Path is the lookup key used in the DNS.
type Path interface {
	GetPrefix() string
	GetContentName() string
	GetBytes() []byte
}
