package impl

import (
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/cmds/agent/sa"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all preregistered implementation IDs for a service agent",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) error {
		err2.Check(cmdList.Validate())
		if !cmd.DryRun() {
			c.SilenceUsage = true
			err2.Try(cmdList.Exec(os.Stdout))
		}
		return nil
	},
}

var cmdList = sa.ListCmd{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	implCmd.AddCommand(listCmd)
}
