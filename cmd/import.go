package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Command for importing wallet",
	Long: `
Command for importing wallet

Example
	findy-agent-cli tools import \
		--wallet-name MyWallet \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--key walletImportKey \
		--file /path/to/my-import-wallet
	`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(impCmd.Validate())
		if !rootFlags.dryRun {
			err2.Try(impCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var impCmd = agent.ImportCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	flags := importCmd.Flags()
	flags.StringVar(&impCmd.WalletName, "wallet-name", "", "wallet name")
	flags.StringVar(&impCmd.WalletKey, "wallet-key", "", "wallet key")
	flags.StringVar(&impCmd.Filename, "file", "", "full import file path")
	flags.StringVar(&impCmd.Key, "key", "", "wallet import key")

	toolsCmd.AddCommand(importCmd)
}
