package cli

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	dnstypes "github.com/AdityaSripal/sangam/dns/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// TODO: add flag to load content by file and then hash
// TODO: add owners flag to extend ownership

var (
	seed   = int64(10)
	random = rand.New(rand.NewSource(seed))
)

// GetCmdPreCommitEntry defines the command to create a pre-commit entry.
func GetCmdPreCommitEntry(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pre-commit [domain] [content_hash]",
		Short: "pre-commit a dns entry with a domain and content hash",
		Long: strings.TrimSpace(fmt.Sprintf(`pre-commit a dns entry with a domain (prefix/content_name) and content hash:

Example:
$ %s tx dns pre-commit [domain] [content_hash] --from node0 --owners [sdk.AccAddress, sdk.AccAddress, sdk.AccAddress]
	`, version.ClientName),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			domain, err := dnstypes.StringToDomain(args[0])
			if err != nil {
				return err
			}

			contentHash := []byte(args[1])
			owners := []sdk.AccAddress{cliCtx.GetFromAddress()}

			// FIXME: use secure random number generator
			nonce := sdk.Uint64ToBigEndian(random.Uint64())
			preCommitEntry := dnstypes.Entry{
				Owners:      owners,
				ContentHash: contentHash,
				Sequence:    dnstypes.PreCommitSequence,
			}
			value := append(nonce, preCommitEntry.GetBytes()...)
			hasher := sha256.New()
			hash := hasher.Sum(value)

			msg := dnstypes.NewMsgPreCommitEntry(domain, hash, owners)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			fmt.Printf("the random nonce used was: %d", nonce)
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}

// GetCmdCommitEntry defines the command to create a commit entry.
func GetCmdCommitEntry(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit [nonce] [domain] [content_hash]",
		Short: "commit a dns entry by submitting the nonce used in a pre-commit",
		Long: strings.TrimSpace(fmt.Sprintf(`commit a dns entry by proving a pre-commit entry through revealing the random nonce used:

Example:
$ %s tx dns commit [nonce] [domain] [content_hash] --from node0 --owners [sdk.AccAddress, sdk.AccAddress, sdk.AccAddress]
	`, version.ClientName),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			nonce, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			domain, err := dnstypes.StringToDomain(args[1])
			if err != nil {
				return err
			}

			contentHash := []byte(args[2])
			owners := []sdk.AccAddress{cliCtx.GetFromAddress()}

			entry := dnstypes.Entry{
				Owners:      owners,
				ContentHash: contentHash,
				Sequence:    dnstypes.CommitSequence,
			}

			msg := dnstypes.NewMsgCommitEntry(nonce, domain, entry, owners)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}
