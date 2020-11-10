package jwt

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/agent/utils"
	"github.com/findy-network/findy-agent/grpc/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var saListenCmd = &cobra.Command{
	Use:   "salisten",
	Short: "SA listen command for JWT gRPC",
	Long: `
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", cmdData.APIService, cmdData.Port, nil)
		conn = client.TryOpen(cmdData.CaDID, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		ch, err := client.Listen(ctx, &agency.ClientID{Id: utils.UUID()})
		err2.Check(err)

	loop:
		for {
			select {
			case status, ok := <-ch:
				if !ok {
					fmt.Println("closed from server")
					break loop
				}
				fmt.Println("listen status:", status.ClientId, status.Notification.TypeId, status.Notification.Id, status.Notification.ProtocolId)
				switch status.Notification.TypeId {
				case agency.Notification_ACTION_NEEDED:
					resume(status, true)
				case agency.Notification_ANSWER_NEEDED_PING:
					reply(status, true)
				case agency.Notification_ANSWER_NEEDED_ISSUE_PROPOSE:
					reply(status, true)
				case agency.Notification_ANSWER_NEEDED_PROOF_PROPOSE:
					reply(status, true)
				case agency.Notification_ANSWER_NEEDED_PROOF_VERIFY:
					reply(status, true)
				}
			case <-intCh:
				cancel()
				fmt.Println("interrupted by user, cancel() called")
			}
		}

		return nil
	},
}

func reply(status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	c := agency.NewAgentClient(conn)
	cid, err := c.Give(ctx, &agency.Answer{
		Id:       status.Notification.Id,
		ClientId: status.ClientId,
		Ack:      ack,
		Info:     "cmd salisten says hello!",
	})
	err2.Check(err)
	fmt.Printf("Sending the answer (%s) send to client:%s\n", status.Notification.Id, cid.Id)
}

func resume(status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewDIDCommClient(conn)
	stateAck := agency.ProtocolState_ACK
	if !ack {
		stateAck = agency.ProtocolState_NACK
	}
	unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
		ProtocolId: &agency.ProtocolID{
			TypeId: agency.Protocol_PROOF,
			Role:   agency.Protocol_RESUME,
			Id:     status.Notification.ProtocolId,
		},
		State: stateAck,
	})
	err2.Check(err)
	fmt.Println("result:", unpauseResult.String())
}

var conn *grpc.ClientConn
var ack bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	jwtCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")
	jwtCmd.AddCommand(saListenCmd)
}
