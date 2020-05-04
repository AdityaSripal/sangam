package keeper

import (
	"github.com/adityasripal/sangam/dns/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper represents a type that reads and writes DNS entries.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// SetPreCommit stores pre-commit information.
func (k Keeper) SetPreCommit(ctx sdk.Context, path types.Path, preCommitHash []byte) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(path.GetBytes(), preCommitHash)
}

// GetPreCommit gets the pre-commit with the given path.
func (k Keeper) GetPreCommit(ctx sdk.Context, path types.Path) (path, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(path.GetBytes())
	if bz == nil {
		return nil, false
	}

	return bz, true
}
