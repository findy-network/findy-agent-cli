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
	Long:  `Long description & example todo`,
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
	flags.StringVar(&msgCmd.Name, "con-id", "", "connection id")
	err2.Check(sendCmd.MarkFlagRequired("msg"))
	err2.Check(sendCmd.MarkFlagRequired("con-id"))

	serviceCopy := *sendCmd
	userCmd.AddCommand(sendCmd)
	serviceCmd.AddCommand(&serviceCopy)
}
