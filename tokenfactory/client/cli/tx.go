package cli

import (
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/Stride-Labs/tokenfactory/tokenfactory/types"
	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := osmocli.TxIndexCmd(types.ModuleName)
	cmd.AddCommand(
		NewCreateDenomCmd(),
		NewMintCmd(),
		NewBurnCmd(),
		NewChangeAdminCmd(),
		NewMsgSetDenomMetadata(),
	)

	return cmd
}

func NewMsgSetDenomMetadata() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgSetDenomMetadata](&osmocli.TxCliDesc{
		Use:   "set-denom-metadata",
		Short: "overwriting of the denom metadata in the bank module.",
	})
}

func NewCreateDenomCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgCreateDenom](&osmocli.TxCliDesc{
		Use:   "create-denom",
		Short: "create a new denom from an account. (osmo to create tokens is charged through gas consumption)",
	})
}

func NewMintCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgMint](&osmocli.TxCliDesc{
		Use:   "mint",
		Short: "Mint a denom to an address. Must have admin authority to do so.",
	})
}

func NewBurnCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgBurn](&osmocli.TxCliDesc{
		Use:   "burn",
		Short: "Burn tokens from an address. Must have admin authority to do so.",
	})
}

func NewChangeAdminCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgChangeAdmin](&osmocli.TxCliDesc{
		Use:   "change-admin",
		Short: "Changes the admin address for a factory-created denom. Must have admin authority to do so.",
	})
}
