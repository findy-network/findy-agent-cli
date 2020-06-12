package cmd

import (
	"log"

	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Parent command for user client",
	Long: `
Parent command for user agent actions.

This command requires a subcommand so command itself does nothing.
Every user subcommand requires --wallet-name & --wallet-key flags to be specified.
--agency-url flag is agency endpoint base address & it has default value of "http://localhost:8080".

Example
	findy-agent-cli user ping \
		--wallet-name TestWallet \
		--wallet-key 6cih1cVgRH8yHD54nEYyPKLmdv67o8QbufxaTHot3Qxp
`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	flags := userCmd.PersistentFlags()
	flags.StringVar(&cFlags.WalletName, "wallet-name", "", "wallet name")
	flags.StringVar(&cFlags.WalletKey, "wallet-key", "", "wallet key")
	flags.StringVar(&cFlags.URL, "agency-url", "http://localhost:8080", "endpoint base address")

	err2.Check(userCmd.MarkPersistentFlagRequired("wallet-name"))
	err2.Check(userCmd.MarkPersistentFlagRequired("wallet-key"))

	rootCmd.AddCommand(userCmd)
}
