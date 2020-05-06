package keeper

import (
	"github.com/AdityaSripal/sangam/dns/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper represents a type that reads and writes DNS entries.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// SetEntry stores an entry at the given domain.
func (k Keeper) SetEntry(ctx sdk.Context, domain types.Domain, entry types.Entry) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&entry)
	store.Set(domain.GetBytes(), bz)
}

// GetEntry gets the entry with the given domain.
func (k Keeper) GetEntry(ctx sdk.Context, domain types.Domain) (entry types.Entry, _ bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(domain.GetBytes())
	if bz == nil {
		return entry, false
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &entry)
	return entry, true
}
