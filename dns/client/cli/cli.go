package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for DNS module commands.
func GetTxCmd(cdc *codec.Codec, storeKey string) *cobra.Command {
	dnsTxCmd := &cobra.Command{
		Use:                        "dns",
		Short:                      "dns transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	dnsTxCmd.AddCommand(flags.PostCommands(
		GetCmdPreCommitEntry(cdc),
		GetCmdCommitEntry(cdc),
	)...)

	return dnsTxCmd
}
