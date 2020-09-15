package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// trustpingCmd represents the trustping subcommand
var trustpingCmd = &cobra.Command{
	Use:   "trustping",
	Short: "Command for making trustping to another agent",
	Long: `
Command for making trustping to another agent

Example
	findy-agent-cli user trustping \
		--wallet-name TheNewWallet4 \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--connection-id my_connection_id
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("connection-id", envPrefix+"TRUSTPING_CONNECTION_ID"))
		return nil
	},
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
	f.StringVar(&tPingCmd.Name, "connection-id", "", "connection id, ENV variable: "+envPrefix+"TRUSTPING_CONNECTION_ID")

	userCmd.AddCommand(trustpingCmd)
	serviceCopy := *trustpingCmd
	serviceCmd.AddCommand(&serviceCopy)
}
