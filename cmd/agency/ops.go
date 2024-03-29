package agency

import (
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// OpsCmd represents the JWT toke based user commands
var OpsCmd = &cobra.Command{
	Use:   "agency",
	Short: "Manage Agency",
	Long:  ``,
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, _ []string) {
		cmd.SubCmdNeeded(c)
	},
}

var CmdData = struct {
	JWT string
}{}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	flags := OpsCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "",
		cmd.FlagInfo("Agency management JWT", "", envs["jwt"]))

	OpsCmd.MarkPersistentFlagRequired("jwt")
	cmd.RootCmd().AddCommand(OpsCmd)
}

var envs = map[string]string{
	"jwt": "JWT",
}
