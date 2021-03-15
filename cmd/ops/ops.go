package ops

import (
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// OpsCmd represents the JWT toke based user commands
var OpsCmd = &cobra.Command{
	Use:   "devops",
	Short: "Operations to manage the cloud agency",
	Long:  ``,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
	},
}

const timeout = 30000 * time.Second

var CmdData = struct {
	JWT string
}{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := OpsCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "", "JWT")

	cmd.RootCmd().AddCommand(OpsCmd)
}

var envs = map[string]string{
	"jwt": "JWT",
}
