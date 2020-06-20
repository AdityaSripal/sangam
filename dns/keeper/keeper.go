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
	if bz == nil {
		return types.Domain{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &dom)
	return dom, true
}

func (k Keeper) SetContent(ctx sdk.Context, content types.Content) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&content)
	key := types.GetCommitKey(content.GetParent(), content.GetName())
	store.Set(key, bz)
}

func (k Keeper) UpdateContent(ctx sdk.Context, domain, name string, latest []byte) (ok bool) {
	content, ok := k.GetContent(ctx, domain, name)
	if !ok {
		return false
	}

	content.ContentHashes = append(content.ContentHashes, latest)
	k.SetContent(ctx, content)
	return true
}

func (k Keeper) GetContent(ctx sdk.Context, domain, name string) (types.Content, bool) {
	store := ctx.KVStore(k.storeKey)
	var content types.Content
	bz := store.Get(types.GetCommitKey(domain, name))
	if bz == nil {
		return types.Content{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &content)
	return content, true
}

func (k Keeper) SetPrecommit(ctx sdk.Context, domain, name string, precommit []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetCommitKey(domain, name)
	store.Set(key, precommit)
}

func (k Keeper) GetPrecommit(ctx sdk.Context, domain, name string) (precommit []byte, ok bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetCommitKey(domain, name)
	bz := store.Get(key)
	if bz == nil {
		return nil, false
	}
	return bz, true
}
