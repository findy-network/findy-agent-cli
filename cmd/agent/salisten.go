package agent

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var saListenDoc = `Starts to cloud agent controller.

The controller runs until interrupted with ctrl-c. During the execution the
controller resumes all protocol steps according the given ACK or NACK flag.`

var saListenCmd = &cobra.Command{
	Use:   "salisten",
	Short: "Start to listen service agent",
	Long:  saListenDoc,
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(*cobra.Command, []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		// first let's ping our agent to get proper error message for
		// JWT authentication and we are not hurry
		timeout, timeoutCancel := context.WithTimeout(
			context.Background(), pingTimeout)
		defer timeoutCancel()

		agent := agency.NewAgentServiceClient(conn)
		try.To1(agent.Ping(timeout, &agency.PingMsg{
			ID:             1000,
			PingController: andController,
		}))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM, syscall.SIGINT)

		ch := try.To1(conn.Listen(ctx, &agency.ClientID{ID: uuid.New().String()}))

	loop:
		for {
			select {
			case question, ok := <-ch:
				if !ok {
					break loop
				}
				status := question.Status
				fmt.Println("listen status:\n",
					"  ClientID:", status.ClientID,
					"  TypeID:", status.Notification.TypeID,
					"  Notification.ID:", status.Notification.ID,
					"  ProtocolID:", status.Notification.ProtocolID,
				)
				switch status.Notification.TypeID {
				case agency.Notification_PROTOCOL_PAUSED:
					resume(status, true)
				}
				switch question.TypeID {
				case agency.Question_PING_WAITS:
					reply(status, true)
				case agency.Question_ISSUE_PROPOSE_WAITS:
					reply(status, true)
				case agency.Question_PROOF_PROPOSE_WAITS:
					reply(status, true)
				case agency.Question_PROOF_VERIFY_WAITS:
					reply(status, true)
				}
			case <-intCh:
				cancel()
			}
		}

		return nil
	},
}

func reply(status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	c := agency.NewAgentServiceClient(conn)
	cid := try.To1(c.Give(ctx, &agency.Answer{
		ID:       status.Notification.ID,
		ClientID: status.ClientID,
		Ack:      ack,
		Info:     "cmd salisten says hello!",
	}))
	fmt.Printf("Sending the answer (%s) send to client:%s\n", status.Notification.ID, cid.ID)
}

func resume(status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewProtocolServiceClient(conn)
	stateAck := agency.ProtocolState_ACK
	if !ack {
		stateAck = agency.ProtocolState_NACK
	}
	unpauseResult := try.To1(didComm.Resume(ctx, &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_PRESENT_PROOF,
			Role:   agency.Protocol_RESUMER,
			ID:     status.Notification.ProtocolID,
		},
		State: stateAck,
	}))
	fmt.Println("result:", unpauseResult.String())
}

var conn client.Conn
var ack bool

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	saListenCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true,
		"used reply ack for all request")

	AgentCmd.AddCommand(saListenCmd)
}
