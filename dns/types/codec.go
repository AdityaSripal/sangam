package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc defines the dns module codec.
var ModuleCdc *codec.Codec

// RegisterCodec registers the Tendermint types
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(Domain{}, "dns/types/Domain", nil)
	cdc.RegisterConcrete(MsgPreCommitEntry{}, "dns/types/MsgPreCommitEntry", nil)
	cdc.RegisterConcrete(MsgCommitEntry{}, "dns/types/MsgCommitEntry", nil)

	SetModuleCodec(cdc)
}

// SetModuleCodec sets the dns module codec
func SetModuleCodec(cdc *codec.Codec) {
	ModuleCdc = cdc
}
