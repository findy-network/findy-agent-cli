package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
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
		--connection-id 1868c791-04a7-4160-bdce-646b975c8de1
		--msg Hello world!
`,
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
	flags.StringVar(&msgCmd.Sender, "from", "", "name of the msg sender")
	flags.StringVar(&msgCmd.Message, "msg", "", "message to be send")
	flags.StringVar(&msgCmd.Name, "connection-id", "", "connection id")
	err2.Check(sendCmd.MarkFlagRequired("msg"))
	err2.Check(sendCmd.MarkFlagRequired("connection-id"))

	serviceCopy := *sendCmd
	userCmd.AddCommand(sendCmd)
	serviceCmd.AddCommand(&serviceCopy)
}
