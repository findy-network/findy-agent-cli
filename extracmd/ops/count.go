package ops

import (
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/cmds/agency"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(cCmd.Validate())
		if !cmd.DryRun() {
			c.SilenceUsage = true
			err2.Try(cCmd.RpcExec(os.Stdout))
		}
		return nil
	},
}

var cCmd agency.CountCmd

func init() {
	countCmd.Flags().StringVar(&cCmd.Addr, "address", "localhost", "gRPC server address")
	countCmd.Flags().IntVar(&cCmd.Port, "port", 50051, "gRPC server port")
	cmd.AgencyCmd.AddCommand(countCmd)
}
