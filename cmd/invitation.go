package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// invitationCmd represents the invitation subcommand
var invitationCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Command for creating invitation message for agent",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		invitateCmd.WalletName = cFlags.WalletName
		invitateCmd.WalletKey = cFlags.WalletKey
		err2.Check(invitateCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(invitateCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var invitateCmd = agent.InvitationCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	invitationCmd.Flags().StringVar(&invitateCmd.Name, "label", "", "invitation label")
	err2.Check(invitationCmd.MarkFlagRequired("label"))

	userCmd.AddCommand(invitationCmd)
	serviceCopy := *invitationCmd
	serviceCmd.AddCommand(&serviceCopy)
}
