package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// exportCmd represents the export subcommand
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Command for exporting wallet",
	Long: `
Command for exporting wallet

Example
	findy-agent-cli tools export \
		--wallet-name MyWallet \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--key walletExportKey \
		--file path/to/my-export-wallet
	`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(expCmd.Validate())
		if !rootFlags.dryRun {
			err2.Try(expCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var expCmd = agent.ExportCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	flags := exportCmd.Flags()
	flags.StringVar(&expCmd.WalletName, "wallet-name", "", "wallet name")
	flags.StringVar(&expCmd.WalletKey, "wallet-key", "", "wallet key")
	flags.StringVar(&expCmd.Filename, "file", "", "full export file path")
	flags.StringVar(&expCmd.ExportKey, "key", "", "wallet export key")

	toolsCmd.AddCommand(exportCmd)
}
