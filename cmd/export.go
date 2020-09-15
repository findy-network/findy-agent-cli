package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("wallet-name", envPrefix+"_EXPORT_WALLET_NAME"))
		err2.Check(viper.BindEnv("wallet-key", envPrefix+"_EXPORT_WALLET_KEY"))
		err2.Check(viper.BindEnv("file", envPrefix+"_EXPORT_WALLET_FILE"))
		err2.Check(viper.BindEnv("key", envPrefix+"_EXPORT_WALLET_FILE_KEY"))
		return nil
	},
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
	flags.StringVar(&expCmd.WalletName, "wallet-name", "", "wallet name, ENV variable: "+envPrefix+"_EXPORT_WALLET_NAME")
	flags.StringVar(&expCmd.WalletKey, "wallet-key", "", "wallet key, ENV variable: "+envPrefix+"_EXPORT_WALLET_KEY")
	flags.StringVar(&expCmd.Filename, "file", "", "full export file path, ENV variable: "+envPrefix+"_EXPORT_WALLET_FILE")
	flags.StringVar(&expCmd.ExportKey, "key", "", "wallet export key, ENV variable: "+envPrefix+"_EXPORT_WALLET_FILE_KEY")

	toolsCmd.AddCommand(exportCmd)
}
