package types

// Path represents the key used to map to a DNS entry. It contains an optional
// prefix. Prefixes are useful in allowing a single entity to register multiple
// content under the same name.
type Path struct {
	Prefix string `json:"prefix" yaml:"prefix"`

	ContentName string `json:"content_name" yaml:"content_name"`
}

// GetPrefix returns the path prefix.
func (p Path) GetPrefix() string {
	return p.Prefix
}

// GetContentName returns the name of the content for the path.
func (p Path) GetContentName() string {
	return p.ContentName
}

// GetBytes returns the path in bytes.
func (p Path) GetBytes() []byte {
	return append([]byte(p.Prefix), []byte(p.ContentName)...)
}
