package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// trustpingCmd represents the trustping subcommand
var trustpingCmd = &cobra.Command{
	Use:   "trustping",
	Short: "Command for making trustping to another agent",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		tPingCmd.WalletName = cFlags.WalletName
		tPingCmd.WalletKey = cFlags.WalletKey
		err2.Check(tPingCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(tPingCmd.Exec(os.Stdout))
		}
		return nil
	},
}
var tPingCmd = connection.TrustPingCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})
	f := trustpingCmd.Flags()
	f.StringVar(&tPingCmd.Name, "con-id", "", "connection id")
	err2.Check(trustpingCmd.MarkFlagRequired("con-id"))

	userCmd.AddCommand(trustpingCmd)
	serviceCopy := *trustpingCmd
	serviceCmd.AddCommand(&serviceCopy)
}
