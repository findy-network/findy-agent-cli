package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/findy-network/findy-agent/cmds/onboard"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// onboardCmd represents the onboard subcommand
var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Command for onboarding agent",
	Long: `
Command for onboarding agent.

If --export-file & --export-key flags are set, 
wallet is exported to that location.
	
Example
	findy-agent-cli user onboard \
		--wallet-name TheNewWallet4 \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp	\
		--email myExampleEmail \
		--salt mySalt
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("export-file", envPrefix+"_ONBOARD_EXPORT_FILE"))
		err2.Check(viper.BindEnv("export-key", envPrefix+"_ONBOARD_EXPORT_KEY"))
		err2.Check(viper.BindEnv("email", envPrefix+"_ONBOARD_EMAIL"))
		err2.Check(viper.BindEnv("salt", envPrefix+"_ONBOARD_SALT"))
		return nil
	},
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
	flags.StringVar(&onbExpCmd.Filename, "export-file", "", "full export file path, ENV variable: "+envPrefix+"_ONBOARD_EXPORT_FILE")
	flags.StringVar(&onbExpCmd.ExportKey, "export-key", "", "wallet export key, ENV variable: "+envPrefix+"_ONBOARD_EXPORT_KEY")
	flags.StringVar(&onbCmd.Email, "email", "", "onboarding email, ENV variable: "+envPrefix+"_ONBOARD_EMAIL")
	flags.StringVar(&aCmd.Salt, "salt", "", "onboarding salt, ENV variable: "+envPrefix+"_ONBOARD_SALT")

	serviceCopy := *onboardCmd
	userCmd.AddCommand(onboardCmd)
	serviceCmd.AddCommand(&serviceCopy)

}
