package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// invitationCmd represents the invitation subcommand
var invitationCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Command for creating invitation message for agent",
	Long: `
Command for creating invitation message for agent	

Example
	findy-agent-cli user invitation \
		--wallet-name MyWallet \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--label invitation_label
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("label", envPrefix+"_INVITATION_LABEL"))
		return nil
	},
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

	invitationCmd.Flags().StringVar(&invitateCmd.Name, "label", "", "invitation label, ENV variable: "+envPrefix+"_INVITATION_LABEL")

	userCmd.AddCommand(invitationCmd)
	serviceCopy := *invitationCmd
	serviceCmd.AddCommand(&serviceCopy)
}
