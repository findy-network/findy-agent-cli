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
	"github.com/spf13/cobra"
)

var saListenDoc = `Starts to cloud agent controller.

The controller runs until interrupted with ctrl-c. During the execution the
controller resumes all protocol steps according the given ACK or NACK flag.`

var saListenCmd = &cobra.Command{
	Use:   "salisten",
	Short: "Start to listen service agent",
	Long:  saListenDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		ch, err := conn.Listen(ctx, &agency.ClientID{ID: uuid.New().String()})
		err2.Check(err)

	loop:
		for {
			select {
			case question, ok := <-ch:
				if !ok {
					fmt.Println("closed from server")
					break loop
				}
				status := question.Status
				fmt.Println("listen status:", status.ClientID, status.Notification.TypeID, status.Notification.ID, status.Notification.ProtocolID)
				switch status.Notification.TypeID {
				case agency.Notification_PROTOCOL_PAUSED:
					resume(status, true)
				}
				switch question.TypeID {
				case agency.Question_ANSWER_NEEDED_PING:
					reply(status, true)
				case agency.Question_ANSWER_NEEDED_ISSUE_PROPOSE:
					reply(status, true)
				case agency.Question_ANSWER_NEEDED_PROOF_PROPOSE:
					reply(status, true)
				case agency.Question_ANSWER_NEEDED_PROOF_VERIFY:
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
	c := agency.NewAgentServiceClient(conn)
	cid, err := c.Give(ctx, &agency.Answer{
		ID:       status.Notification.ID,
		ClientID: status.ClientID,
		Ack:      ack,
		Info:     "cmd salisten says hello!",
	})
	err2.Check(err)
	fmt.Printf("Sending the answer (%s) send to client:%s\n", status.Notification.ID, cid.ID)
}

func resume(status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewProtocolServiceClient(conn)
	stateAck := agency.ProtocolState_ACK
	if !ack {
		stateAck = agency.ProtocolState_NACK
	}
	unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_PRESENT_PROOF,
			Role:   agency.Protocol_RESUMER,
			ID:     status.Notification.ProtocolID,
		},
		State: stateAck,
	})
	err2.Check(err)
	fmt.Println("result:", unpauseResult.String())
}

var conn client.Conn
var ack bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	saListenCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")

	AgentCmd.AddCommand(saListenCmd)
}
