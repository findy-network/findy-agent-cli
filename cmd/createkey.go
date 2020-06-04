package cmd

import (
	"log"
	"os"

	"github.com/lainio/err2"
	"github.com/optechlab/findy-agent/cmds/key"
	"github.com/spf13/cobra"
)

// createKeyCmd represents the createkey subcommand
var createKeyCmd = &cobra.Command{
	Use:   "createkey",
	Short: "Command for creating valid wallet keys",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(keyCreateCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(keyCreateCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var keyCreateCmd = key.CreateCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	createKeyCmd.Flags().StringVar(&keyCreateCmd.Seed, "seed", "", "Seed for wallet key creation")
	err2.Check(createKeyCmd.MarkFlagRequired("seed"))

	userCmd.AddCommand(createKeyCmd)
	serviceCopy := *createKeyCmd
	serviceCmd.AddCommand(&serviceCopy)
}
