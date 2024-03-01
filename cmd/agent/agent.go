package agent

import (
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// AgentCmd represents the JWT toke based user commands
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Work with the Cloud Agent",
	Long:  ``,
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, _ []string) {
		cmd.SubCmdNeeded(c)
	},
}

const timeout = 30000 * time.Second

var CmdData = struct {
	JWT string
}{}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	flags := AgentCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "",
		cmd.FlagInfo("Agent JWT token", "", envs["jwt"]))

	cmd.RootCmd().AddCommand(AgentCmd)
}

var envs = map[string]string{
	"jwt": "JWT",
}
