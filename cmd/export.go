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
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		expCmd.WalletName = cFlags.WalletName
		expCmd.WalletKey = cFlags.WalletKey
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
	flags.StringVar(&expCmd.Filename, "file", "", "full export file path")
	flags.StringVar(&expCmd.ExportKey, "key", "", "wallet export key")
	err2.Check(exportCmd.MarkFlagRequired("file"))
	err2.Check(exportCmd.MarkFlagRequired("key"))

	userCmd.AddCommand(exportCmd)
	serviceCopy := *exportCmd
	serviceCmd.AddCommand(&serviceCopy)
}
