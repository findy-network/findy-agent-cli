package connection

import (
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/ghodss/yaml"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

// ConnectionCmd represents the JWT toke based user commands
var ConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Manage connections",
	Long:  `Manages DIDComm based protocol communication.`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
	},
}

const timeout = 30000 * time.Second

var CmdData = struct {
	ConnID string
	JWT    string
}{}

func PrintCmdData() {
	cb := try.To1(yaml.Marshal(CmdData))
	fmt.Println(string(cb))
}

var conn client.Conn

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := ConnectionCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "",
		cmd.FlagInfo("cloud agent JWT token", "", envs["jwt"]))
	flags.StringVar(&CmdData.ConnID, "conn-id", "",
		cmd.FlagInfo("connection id aka pairwise id", "", envs["conn-id"]))

	ConnectionCmd.MarkPersistentFlagRequired("jwt")
	ConnectionCmd.MarkPersistentFlagRequired("conn-id")
	cmd.RootCmd().AddCommand(ConnectionCmd)
}

var envs = map[string]string{
	"jwt":     "JWT",
	"conn-id": "CONN_ID",
}
