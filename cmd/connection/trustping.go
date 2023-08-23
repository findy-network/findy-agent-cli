package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "trustping",
	Short: "Trustping protocol",
	Long:  `Executes Aries trust ping protocol.`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: trustping,
}

func trustping(_ *cobra.Command, _ []string) (err error) {
	defer err2.Handle(&err)

	if cmd.DryRun() {
		PrintCmdData()
		return nil
	}
	baseCfg := try.To1(cmd.BaseCfg())
	conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := try.To1(client.Pairwise{ID: CmdData.ConnID, Conn: conn}.Ping(ctx))

	okOutput := false
	for status := range ch {
		if status.State == agency.ProtocolState_ERR {
			fmt.Fprintln(os.Stderr, "protocol error:", status.GetInfo())
			err = fmt.Errorf("protocol error: %s", status.GetInfo())
		} else {
			okOutput = true
			fmt.Println("ping status:", status.State)
		}
	}
	if !okOutput {
		err = fmt.Errorf("no response")
	}
	return err
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))
	ConnectionCmd.AddCommand(pingCmd)
}
