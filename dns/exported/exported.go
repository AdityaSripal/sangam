package exported

// Entry is the entry used in the DNS.
type Entry interface {
	GetOwners() [][]byte
	GetContent() []byte
}

// MsgPreCommitEntry defines the msg interface used to send a pre-commit of a
// registration in the DNS.
type MsgPreCommitEntry interface {
	GetKey() []byte
	GetHash() []byte
}

// MsgCommitEntry defines the msg interface used to send a commit of a registration
// in the DNS.
type MsgCommitEntry interface {
	GetNonce() uint64
	GetKey() []byte
	GetEntry() Entry
}
