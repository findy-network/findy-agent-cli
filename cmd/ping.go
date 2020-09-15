package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pingCmd represents the user/service ping subcommand
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Command for pinging services and agents",
	Long: ` 
Tests the connection to the CA with the given wallet. If secure connection works
ok it prints the invitation. If the EA is a SA the command pings it as well when
the --service-endpoint flag is on.

Example
	findy-agent-cli user ping \
		--service-endpoint \
		--wallet-name TheNewWallet4 \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp

	this pings the CA and the connected SA as well. 
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("service-endpoint", envPrefix+"_PING_SERVICE_ENDPOINT"))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		pCmd.WalletName = cFlags.WalletName
		pCmd.WalletKey = cFlags.WalletKey
		err2.Check(pCmd.Validate())
		if !rootFlags.dryRun {
			// if error occurs in the execution, we don't show usage, only
			// the error message.
			cmd.SilenceUsage = true

			r, err := pCmd.Exec(os.Stdout)
			err2.Check(err)
			jBytes := err2.Bytes.Try(r.JSON())
			fmt.Println(string(jBytes))
		}
		return nil
	},
}

var pCmd = agent.PingCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	pingCmd.Flags().BoolVarP(&pCmd.PingSA, "service-endpoint", "s", false, "ping CA and connected SA (me) as well, ENV variable: "+envPrefix+"_PING_SERVICE_ENDPOINT")

	// service copy
	serviceCopy := *pingCmd
	userCmd.AddCommand(pingCmd)
	serviceCmd.AddCommand(&serviceCopy)
}
