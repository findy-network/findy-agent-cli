package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
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
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(*cobra.Command, []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			fmt.Printf("read: %v\n", read)
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentServiceClient(conn)
		mode := agency.ModeCmd_AcceptModeCmd_GRPC_CONTROL
		if auto {
			mode = agency.ModeCmd_AcceptModeCmd_AUTO_ACCEPT
		}
		r := try.To1(agent.Enter(ctx, &agency.ModeCmd{
			TypeID:  agency.ModeCmd_ACCEPT_MODE,
			IsInput: !read,
			ControlCmd: &agency.ModeCmd_AcceptMode{
				AcceptMode: &agency.ModeCmd_AcceptModeCmd{
					Mode: mode,
				},
			},
		}))
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
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))
	enterModeCmd.Flags().BoolVarP(&auto, "auto", "a", false,
		"set controller communication mode to auto")
	enterModeCmd.Flags().BoolVarP(&read, "read", "r", false,
		"tells to read communication mode from CA")
	AgentCmd.AddCommand(enterModeCmd)
}
