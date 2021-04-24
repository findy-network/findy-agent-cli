package bot

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	agency "github.com/findy-network/findy-agent-api/grpc/agency/v1"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var readDoc = `The command simulates the chat window input stream where we can hear
what other end sends to us.

The command runs as long as it's stopped with ctrl-c. During the run it echoes 
all of the Aries Basic Message protocol messages to standard output. It also
replies all of the ctrl messages either by ACK or NACK depending of the flag
given at the start.`

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read basic message stream and reply protocol ctrl questions",
	Long:  readDoc,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
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
				glog.V(1).Infoln("listen status:",
					status.ClientID,
					status.Notification.TypeID,
					status.Notification.ID,
					status.Notification.ProtocolID)
				switch status.Notification.TypeID {
				case agency.Notification_STATUS_UPDATE:
					handleBM(conn, status, true)
				case agency.Notification_ACTION_NEEDED:
					resume(conn.ClientConn, status, true)
				}
				switch question.TypeID {
				case agency.Question_ANSWER_NEEDED_PING:
					reply(conn.ClientConn, status, true)
				case agency.Question_ANSWER_NEEDED_ISSUE_PROPOSE:
					reply(conn.ClientConn, status, true)
				case agency.Question_ANSWER_NEEDED_PROOF_PROPOSE:
					reply(conn.ClientConn, status, true)
				case agency.Question_ANSWER_NEEDED_PROOF_VERIFY:
					reply(conn.ClientConn, status, true)
				}
			case <-intCh:
				cancel()
				fmt.Println("interrupted by user, cancel() called")
			}
		}

		return nil
	},
}

var ack bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	readCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")
	botCmd.AddCommand(readCmd)
}

func handleBM(conn client.Conn, status *agency.AgentStatus, _ bool) {
	if status.Notification.ProtocolType == agency.Protocol_BASIC_MESSAGE {
		ctx := context.Background()
		didComm := agency.NewProtocolServiceClient(conn)
		statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
			TypeID:           status.Notification.ProtocolType,
			Role:             agency.Protocol_ADDRESSEE,
			ID:               status.Notification.ProtocolID,
			NotificationTime: status.Notification.Timestamp,
		})
		err2.Check(err)
		if statusResult.GetBasicMessage().SentByMe {
			glog.V(1).Infoln("-- ours, no reply")
			return
		}
		fmt.Println(statusResult.GetBasicMessage().Content)
	}
}

func reply(cc grpc.ClientConnInterface, status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	c := agency.NewAgentServiceClient(cc)
	cid, err := c.Give(ctx, &agency.Answer{
		ID:       status.Notification.ID,
		ClientID: status.ClientID,
		Ack:      ack,
		Info:     "testing says hello!",
	})
	err2.Check(err)
	glog.Infof("Sending the answer (%s) send to client:%s\n", status.Notification.ID, cid.ID)
}

func resume(cc grpc.ClientConnInterface, status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewProtocolServiceClient(cc)
	statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
		TypeID:           status.Notification.ProtocolType,
		ID:               status.Notification.ProtocolID,
		NotificationTime: status.Notification.Timestamp,
	})
	err2.Check(err)
	fmt.Println("** protocol state:", statusResult.State.State)

	stateAck := agency.ProtocolState_ACK
	if !ack {
		stateAck = agency.ProtocolState_NACK
	}
	unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: status.Notification.ProtocolType,
			Role:   agency.Protocol_RESUMER,
			ID:     status.Notification.ProtocolID,
		},
		State: stateAck,
	})
	err2.Check(err)
	glog.Infoln("result:", unpauseResult.String())
}
