package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Domain represents the key used to map to a DNS entry. It contains an optional
// prefix. Prefixes are useful in allowing a single entity to register multiple
// content under the same name.
type Domain struct {
	Prefix string `json:"prefix" yaml:"prefix"`

	ContentName string `json:"content_name" yaml:"content_name"`
}

// GetPrefix returns the domain prefix.
func (d Domain) GetPrefix() string {
	return d.Prefix
}

// GetContentName returns the name of the content for the domain.
func (d Domain) GetContentName() string {
	return d.ContentName
}

// GetBytes returns the domain in bytes.
func (d Domain) GetBytes() []byte {
	return append([]byte(d.Prefix), []byte(d.ContentName)...)
}

// ValidateBasic does basic validation on a Domain object.
// TODO: add stricter checks
func (d Domain) ValidateBasic() error {
	if strings.TrimSpace(d.ContentName) == "" {
		return sdkerrors.Wrap(ErrInvalidContentName, "content name cannot be empty")
	}

	return nil
}

// StringToDomain is a helper function creating a Domain object from a string.
func StringToDomain(domain string) (Domain, error) {
	s := strings.SplitAfterN(domain, "/", 2)
	if len(s) != 2 {
		return Domain{}, sdkerrors.Wrap(ErrInvalidDomainString, "domain string must be in the format `<prefix>/<content_name>`")
	}

	return Domain{
		Prefix:      s[0],
		ContentName: s[1],
	}, nil

}
