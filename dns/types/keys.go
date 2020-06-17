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
	DomainKey = []byte{0x11}
)

func GetDomainKey(domain string) []byte {
	return append(DomainKey, []byte(domain)...)
}
