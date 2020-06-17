package types

import (
	"github.com/AdityaSripal/sangam/dns/exported"
)

// Domain
type Domain struct {
	Owner      exported.DomainOwner `json:"owner" yaml:"owner"`
	SubDomains []string             `json:"sub_domains" yaml:"sub_domains"`
	Contents   []string             `json:"contents" yaml:"contents"`
	Parent     string               `json:"parent" yaml:"parent"`
	Name       string               `json:"name" yaml:"name"`
}

// GetOwner returns the domain owner.
func (d Domain) GetOwner() exported.DomainOwner {
	return d.Owner
}

// GetSubDomains returns the sub-domains underneath this domain entry.
// Subdomains returns just the string of the subdomain that will be appended under the domain's path
func (d Domain) GetSubDomains() []string {
	return d.SubDomains
}

// GetContents returns the contents under this domain entry.
// Contents returns just the part of content's path after this domain's path.
func (d Domain) GetContents() []string {
	return d.Contents
}

// GetParent returns the parent entry for this domain entry.
// Parent will return the full path of the parent domain.
func (d Domain) GetParent() string {
	return d.Parent
}

// String returns the string identifier for this domain.
// Does not include parent path.
func (d Domain) String() string {
	return d.Name
}

// Path returns the full path string representation up to the global top-level
func (d Domain) Path() string {
	return d.Parent + d.Name
}
