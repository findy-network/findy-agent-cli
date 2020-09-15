package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendCmd represents the send subcommand
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Command for sending basic message to another agent",
	Long: `
Sends basic message to another agent.

--msg (message) & --connection-id (id of the connection) flags are required flags on the command.
--from (name of the sender) flag is optional. --connection-id is uuid that is created during agent connection.

Example
	findy-agent-cli user send \
		--wallet-name TestWallet \
		--wallet-key 6cih1cVgRH8yHD54nEYyPKLmdv67o8QbufxaTHot3Qxp \
		--connection-id 1868c791-04a7-4160-bdce-646b975c8de1 \
		--msg Hello world!
`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("from", envPrefix+"_SEND_FROM"))
		err2.Check(viper.BindEnv("msg", envPrefix+"_SEND_MESSAGE"))
		err2.Check(viper.BindEnv("connection-id", envPrefix+"_SEND_CONNECTION_ID"))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		msgCmd.WalletName = cFlags.WalletName
		msgCmd.WalletKey = cFlags.WalletKey
		err2.Check(msgCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(msgCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var msgCmd = connection.BasicMsgCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})
	flags := sendCmd.Flags()
	flags.StringVar(&msgCmd.Sender, "from", "", "name of the msg sender, ENV variable: "+envPrefix+"_SEND_FROM")
	flags.StringVar(&msgCmd.Message, "msg", "", "message to be send, ENV variable: "+envPrefix+"_SEND_MESSAGE")
	flags.StringVar(&msgCmd.Name, "connection-id", "", "connection id, ENV variable: "+envPrefix+"_SEND_CONNECTION_ID")

	serviceCopy := *sendCmd
	userCmd.AddCommand(sendCmd)
	serviceCmd.AddCommand(&serviceCopy)
}
