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
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		impCmd.WalletName = cFlags.WalletName
		impCmd.WalletKey = cFlags.WalletKey
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
	flags.StringVar(&impCmd.Filename, "file", "", "full import file path")
	flags.StringVar(&impCmd.Key, "key", "", "wallet import key")

	userCmd.AddCommand(importCmd)
	serviceCopy := *importCmd
	serviceCmd.AddCommand(&serviceCopy)

}
