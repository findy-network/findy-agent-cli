package jwt

import (
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// JwtCmd represents the JWT toke based user commands
var JwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Parent command for JWT gRPC commands",
	Long: `
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
	},
}

const timeout = 30000 * time.Second

var CmdData = struct {
	APIService string
	Port       int
	ConnID     string
	CaDID      string
	JWT        string
}{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := JwtCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "", "JWT")
	flags.StringVar(&CmdData.CaDID, "ca-did", "", "CA DID")
	flags.StringVar(&CmdData.ConnID, "conn-id", "", "connection id aka pairwise id")
	flags.StringVar(&CmdData.APIService, "server", "localhost", "gRPC server host name")
	flags.IntVar(&CmdData.Port, "port", 50051, "port for gRPC server")

	cmd.RootCmd().AddCommand(JwtCmd)
}

var envs = map[string]string{
	"jwt":     "JWT",
	"ca-did":  "CA_DID",
	"server":  "SERVER",
	"port":    "PORT",
	"conn-id": "CONN_ID",
}
