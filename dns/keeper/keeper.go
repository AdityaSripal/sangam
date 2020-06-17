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

func (k Keeper) SetDomain(ctx sdk.Context, domain types.Domain) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&domain)
	key := types.GetDomainKey(domain.Path())
	store.Set(key, bz)
}

func (k Keeper) GetDomain(ctx sdk.Context, path string) (types.Domain, bool) {
	store := ctx.KVStore(k.storeKey)
	var dom types.Domain
	bz := store.Get(types.GetDomainKey(path))
	k.cdc.MustUnmarshalBinaryBare(bz, &dom)
	return dom, true
}
