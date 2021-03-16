package connection

import (
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// ConnectionCmd represents the JWT toke based user commands
var ConnectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "Manage connections",
	Long:  `Commands for DIDComm based protocol communication`,
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

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := ConnectionCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "", "JWT")
	flags.StringVar(&CmdData.ConnID, "conn-id", "", "connection id aka pairwise id")

	cmd.RootCmd().AddCommand(ConnectionCmd)
}

var envs = map[string]string{
	"jwt":     "JWT",
	"conn-id": "CONN_ID",
}
