package ops

import (
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/cmds/agency"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(lCmd.Validate())
		if !cmd.DryRun() {
			c.SilenceUsage = true
			err2.Try(lCmd.RpcExec(os.Stdout))
		}
		return nil
	},
}

var lCmd agency.LoggingCmd

func init() {
	loggingCmd.Flags().StringVarP(&lCmd.Level, "level", "L", "3", "log level in the agency")
	loggingCmd.Flags().StringVar(&lCmd.BaseAddr, "base-address", "localhost:50051", "gRPC address with port")
	cmd.AgencyCmd.AddCommand(loggingCmd)
}
