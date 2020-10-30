package jwt

import (
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var jwtCmd = &cobra.Command{
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

var cmdData = struct {
	APIService string
	Port       int
	ConnID     string
	CaDID      string
}{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := jwtCmd.PersistentFlags()
	flags.StringVar(&cmdData.CaDID, "ca-did", "", "CA DID")
	flags.StringVar(&cmdData.ConnID, "conn-id", "", "connection id aka pairwise id")
	flags.StringVar(&cmdData.APIService, "server", "localhost", "gRPC server host name")
	flags.IntVar(&cmdData.Port, "port", 50051, "port for gRPC server")

	cmd.RootCmd().AddCommand(jwtCmd)
}

var envs = map[string]string{
	"ca-did":  "CA_DID",
	"server":  "SERVER",
	"port":    "PORT",
	"conn-id": "CONN_ID",
}
