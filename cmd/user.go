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
	Long:  `Long description & example todo`,
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
	flags.StringVar(&cFlags.URL, "url", "http://localhost:8080", "endpoint base address")

	err2.Check(userCmd.MarkPersistentFlagRequired("wallet-name"))
	err2.Check(userCmd.MarkPersistentFlagRequired("wallet-key"))

	rootCmd.AddCommand(userCmd)
}
