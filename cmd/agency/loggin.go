package agency

import (
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/findy-network/findy-agent-api/grpc/ops"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "Manage logging level of the agency",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		if !cmd.DryRun() {
			c.SilenceUsage = true
			err2.Try(Logging(os.Stdout))
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
	defer err2.Return(&err)

	baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
	// todo: this wont work until we have way to build JWT
	conn := client.TryOpen("findy-root", baseCfg)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opsClient := pb.NewDevOpsClient(conn)
	err2.Empty.Try(opsClient.Enter(ctx, &pb.Cmd{
		Type:    pb.Cmd_LOGGING,
		Request: &pb.Cmd_Logging{Logging: lCmd.Level},
	}))
	err2.Check(err)

	fmt.Fprintln(w, "logging level set to:", lCmd.Level)

	return nil
}
