package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AdityaSripal/sangam/dns/exported"
)

// Domain
type Domain struct {
	Owner      exported.DomainOwner `json:"owner" yaml:"owner"`
	SubDomains []Domain             `json:"sub_domains" yaml:"sub_domains"`
	Contents   []Content            `json:"contents" yaml:"contents"`
	Parent     Domain               `json:"parent" yaml:"parent"`
	Name       string               `json:"name" yaml:"name"`
}

// GetOwner returns the domain owner.
func (d Domain) GetOwner() exported.DomainOwner {
	return d.Owner
}

// GetSubDomains returns the sub-domains underneath this domain entry.
func (d Domain) GetSubDomains() []Domain {
	return d.SubDomains
}

// GetContents returns the contents represented at this domain entry.
func (d Domain) GetContents() []Content {
	return d.Contents
}

// GetParent returns the parent entry for this domain entry.
func (d Domain) GetParents() Domain {
	return d.Parent
}

// String returns the string identifier for this domain.
func (d Domain) String() string {
	return d.Name
}

// Path returns the full path string representation up to the global top-level
func (d Domain) Path() string {
	return d.Parent.Path() + d.String()
}

// GetSubDomain returns the subdomain under this domain as specified by the given string.
func (d Domain) GetSubDomain(path string) Domain {
	// TODO: implement
	return Domain{}
}

// GetContent returns the content under this domain as specified by the path. The content
// may either be in this domain or in a sub-domain. The latest version is returned by default.
func (d Domain) GetContent(path string) Content {
	// TODO: implement
	return Domain{}
}

// GetContentAtSequence returns the content at the given sequence.
func (d Domain) GetContentAtSequence(path string, seq uint64) Content {
	return d.Contents[seq]
}

// AddSubDomain adds a sub-domain to the current domain entry.
func (d Domain) AddSubDomain(domain Domain) error {
	if d.GetSubDomain(domain) != nil {
		return sdkerrors.Wrapf(
			ErrEntryExists,
			"cannot add sub-domain (%v) that already exists within this domain (%v)", domain, d,
		)
	}

	d.SubDomains = append(d.SubDomains, domain)
	return nil
}

// AddPreCommit adds a pre-commit related to this domain entry.
func (d Domain) AddPreCommit(name string, precommit []byte) error {
	// TODO: implement
	return nil
}

// AddContent adds a content entry to this domain.
func (d Domain) AddContent(content Content, reveal uint64) error {
	// TODO: implement

	if d.GetContent(content) != nil {
		return sdkerrors.Wrapf(
			ErrEntryExists,
			"cannot add content (%v) that already exists within this domain (%v)", content, d,
		)
	}

	d.Contents = append(d.Contents, content)
	return nil
}

// UpdateSubDomain updates the sub-domain contained within this entry.
func (d Domain) UpdateSubDomain(name string, domain Domain) error {
	for i, subDomain := range d.SubDomains {
		if name == subDomain.GetName() {
			d.SubDomains[i] = domain
			return nil
		}
	}
	return sdkerrors.Wrapf(ErrDomainNotFound, "sub-domain with name (%s) not found for parent domain (%v)", name, d)
}

// UpdateContent updates the content contained within this entry.
// XXX: don't you need a sequence specifiying which content?
func (d Domain) UpdateContent(newHash []byte) error {
	// TODO: implement
	return nil
}

// DeleteSubDomain deletes the subdomain with the provided name.
func (d Domain) DeleteSubDomain(name string) error {
	for i, subDomain := range d.SubDomains {
		if name == subDomain.GetName() {
			d.SubDomains[i] = nil
			return nil
		}
	}
	return sdkerrors.Wrapf(ErrDomainNotFound, "sub-domain with name (%s) not found for parent domain (%v)", name, d)
}

// DeleteContent deletes the content with the provided name under this sub-domain.
func (d Domain) DeleteContent(name string) error {
	for i, content := range d.Contents {
		if name == content.GetName() {
			d.Contents[i] = nil
			return nil
		}
	}
	return sdkerrors.Wrapf(ErrContentNotFound, "content with name (%s) not found for parent domain (%v)", name, d)
}
