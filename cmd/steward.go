package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/steward"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// stewardCmd represents the steward command
var stewardCmd = &cobra.Command{
	Use:   "steward",
	Short: "Parent command for steward wallet actions",
	Long:  `Long description & example todo`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// stewardCreateCmd represents the steward create subcommand
var stewardCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Command for creating steward wallet",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		createStewardCmd.Cmd.WalletName = cFlags.WalletName
		createStewardCmd.Cmd.WalletKey = cFlags.WalletKey
		err2.Check(createStewardCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(createStewardCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var createStewardCmd = steward.CreateCmd{}

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	f := stewardCreateCmd.Flags()
	f.StringVar(&createStewardCmd.PoolName, "pool-name", "FINDY_MEM_LEDGER", "pool name")
	f.StringVar(&createStewardCmd.StewardSeed, "steward-seed", "000000000000000000000000Steward2", "steward seed")

	stewardCmd.AddCommand(stewardCreateCmd)
	ledgerCmd.AddCommand(stewardCmd)
}
