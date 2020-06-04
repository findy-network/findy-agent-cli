package cmd

import (
	"log"
	"os"

	"github.com/lainio/err2"
	"github.com/optechlab/findy-agent/cmds/agent"
	"github.com/optechlab/findy-agent/cmds/onboard"
	"github.com/spf13/cobra"
)

// onboardCmd represents the onboard subcommand
var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Command for onboarding agent",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		onbCmd.WalletName = cFlags.WalletName
		onbCmd.WalletKey = cFlags.WalletKey
		onbCmd.AgencyAddr = cFlags.URL
		err2.Check(onbCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(onbCmd.Exec(os.Stdout))
		}

		if onbExpCmd.Filename != "" {
			onbExpCmd.WalletName = cFlags.WalletName
			onbExpCmd.WalletKey = cFlags.WalletKey
			onbExpCmd.ExportKey = cFlags.WalletKey
			err2.Check(onbExpCmd.Validate())
			if !rootFlags.dryRun {
				err2.Try(onbExpCmd.Exec(os.Stdout))
			}
		}
		return nil
	},
}

var onbCmd = onboard.Cmd{}
var onbExpCmd = agent.ExportCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	flags := onboardCmd.Flags()
	flags.StringVar(&onbExpCmd.Filename, "export-file", "", "filename for wallet export with path")
	flags.StringVar(&onbCmd.Email, "email", "", "onboarding email")

	err2.Check(onboardCmd.MarkFlagRequired("email"))

	serviceCopy := *onboardCmd
	userCmd.AddCommand(onboardCmd)
	serviceCmd.AddCommand(&serviceCopy)

}
