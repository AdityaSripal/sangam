package exported

// Entry is the entry used in the DNS.
type Entry interface {
	GetOwners() [][]byte
	GetContent() []byte
}

// Path is the lookup key used in the DNS.
type Path interface {
	GetPrefix() string
	GetContentName() string
	GetBytes() []byte
}
