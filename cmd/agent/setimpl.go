package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var enterModeDoc = `The command sets a running mode of the cloud agent.

The running mode stands for the protocol which is used for the communication
between cloud agent and its controller. Because we are still in the middle of
the transition to gRPC API we has to support some of the legacy modes.

Note, default mode is GRPC.
`

var enterModeCmd = &cobra.Command{
	Use:   "mode-cmd",
	Short: "Enters the communication mode between CA and its controller",
	Long:  enterModeDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Printf("read: %v\n", read)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentServiceClient(conn)
		mode := agency.ModeCmd_AcceptModeCmd_GRPC_CONTROL
		if auto {
			mode = agency.ModeCmd_AcceptModeCmd_AUTO_ACCEPT
		}
		r, err := agent.Enter(ctx, &agency.ModeCmd{
			TypeID:  agency.ModeCmd_ACCEPT_MODE,
			IsInput: !read,
			ControlCmd: &agency.ModeCmd_AcceptMode{
				AcceptMode: &agency.ModeCmd_AcceptModeCmd{
					Mode: mode,
				},
			},
		})
		mode = r.GetAcceptMode().Mode
		fmt.Print("Current mode:", mode)
		if mode == agency.ModeCmd_AcceptModeCmd_DEFAULT {
			fmt.Println(" default mode is:",
				agency.ModeCmd_AcceptModeCmd_GRPC_CONTROL)
		} else {
			fmt.Println()
		}

		return nil
	},
}

var auto bool
var read bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	enterModeCmd.Flags().BoolVarP(&auto, "auto", "a", false,
		"set controller communication mode to auto")
	enterModeCmd.Flags().BoolVarP(&read, "read", "r", false,
		"tells to read communication mode from CA")
	AgentCmd.AddCommand(enterModeCmd)
}
