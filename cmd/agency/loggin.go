package agency

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	pb "github.com/findy-network/findy-common-go/grpc/ops/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "Manage logging level of the agency",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)
		if !cmd.DryRun() {
			c.SilenceUsage = true
			try.To(Logging(os.Stdout))
		}
		return nil
	},
}

var lCmd struct {
	Level string
}

func init() {
	loggingCmd.Flags().StringVarP(&lCmd.Level, "level", "L", "3", "log level in the agency")
	OpsCmd.AddCommand(loggingCmd)
}

func Logging(w io.Writer) (err error) {
	defer err2.Handle(&err)

	baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
	conn := client.TryAuthOpen(CmdData.JWT, baseCfg)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opsClient := pb.NewDevOpsServiceClient(conn)
	try.To1(opsClient.Enter(ctx, &pb.Cmd{
		Type:    pb.Cmd_LOGGING,
		Request: &pb.Cmd_Logging{Logging: lCmd.Level},
	}))

	fmt.Fprintln(w, "logging level set to:", lCmd.Level)

	return nil
}
