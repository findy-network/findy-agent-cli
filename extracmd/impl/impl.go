package impl

import (
	"fmt"

	"github.com/lainio/err2"
	"github.com/spf13/cobra"

	"github.com/findy-network/findy-agent-cli/cmd"
)

var implCmd = &cobra.Command{
	Use:   "impl",
	Short: "Parent command for service agent implementations",
	Long: `
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		// todo: have to use cmdName until env var names are fixed, now we use "user" could use "service"
		return cmd.BindEnvs(envs, "user")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	cmd.RootCmd().AddCommand(implCmd)
}

var envs = map[string]string{
	"wallet-name": "WALLET_NAME",
	"wallet-key":  "WALLET_KEY",
}
