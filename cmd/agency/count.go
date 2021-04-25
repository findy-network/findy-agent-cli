package agency

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	pb "github.com/findy-network/findy-common-go/grpc/ops/v1"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Query statistics from the agency",
	Long:  ``,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		if !cmd.DryRun() {
			c.SilenceUsage = true
			err2.Try(Count(os.Stdout))
		} else {
			fmt.Println("jwt:", CmdData.JWT)
		}
		return nil
	},
}

const timeout = 10 * time.Second

func init() {
	OpsCmd.AddCommand(countCmd)
}

func Count(w io.Writer) (err error) {
	defer err2.Return(&err)

	baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
	conn := client.TryAuthOpen(CmdData.JWT, baseCfg)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opsClient := pb.NewDevOpsServiceClient(conn)
	result, err := opsClient.Enter(ctx, &pb.Cmd{
		Type: pb.Cmd_COUNT,
	})
	err2.Check(err)
	fmt.Fprintln(w, "count result:\n", result.GetCount())

	return nil
}
