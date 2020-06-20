package types

const (
	// ModuleName defines the IBC transfer name
	ModuleName = "dns"

	// StoreKey is the store key string for IBC transfer
	StoreKey = ModuleName

	// RouterKey is the message route for IBC transfer
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = ModuleName
)

var (
	// Key for store prefixes
	DomainKey    = []byte{0x11}
	PrecommitKey = []byte{0x12}
	CommitKey    = []byte{0x13}
)

func GetDomainKey(domain string) []byte {
	return append(DomainKey, []byte(domain)...)
}

func GetPrecommitKey(domain, name string) []byte {
	return append(PrecommitKey, []byte(domain+name)...)
}

func GetCommitKey(domain, name string) []byte {
	return append(CommitKey, []byte(domain+name)...)
}
