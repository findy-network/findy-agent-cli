package bot

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-agent/agent/utils"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/golang/glog"
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
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", jwt.CmdData.APIService, jwt.CmdData.Port, nil)
		conn = client.TryAuthOpen(jwt.CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		err2.Check(flag.Set("v", "0"))

		ch, err := conn.Listen(ctx, &agency.ClientID{Id: utils.UUID()})
		err2.Check(err)

	loop:
		for {
			select {
			case status, ok := <-ch:
				if !ok {
					fmt.Println("closed from server")
					break loop
				}
				glog.V(1).Infoln("listen status:",
					status.ClientId,
					status.Notification.TypeId,
					status.Notification.Id,
					status.Notification.ProtocolId)
				switch status.Notification.TypeId {
				case agency.Notification_STATUS_UPDATE:
					handleBM(conn, status, true)
				case agency.Notification_ACTION_NEEDED:
					resume(conn.ClientConn, status, true)
				case agency.Notification_ANSWER_NEEDED_PING:
					reply(conn.ClientConn, status, true)
				case agency.Notification_ANSWER_NEEDED_ISSUE_PROPOSE:
					reply(conn.ClientConn, status, true)
				case agency.Notification_ANSWER_NEEDED_PROOF_PROPOSE:
					reply(conn.ClientConn, status, true)
				case agency.Notification_ANSWER_NEEDED_PROOF_VERIFY:
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
		didComm := agency.NewDIDCommClient(conn)
		statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
			TypeId:           status.Notification.ProtocolType,
			Role:             agency.Protocol_ADDRESSEE,
			Id:               status.Notification.ProtocolId,
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
	c := agency.NewAgentClient(cc)
	cid, err := c.Give(ctx, &agency.Answer{
		Id:       status.Notification.Id,
		ClientId: status.ClientId,
		Ack:      ack,
		Info:     "testing says hello!",
	})
	err2.Check(err)
	glog.Infof("Sending the answer (%s) send to client:%s\n", status.Notification.Id, cid.Id)
}

func resume(cc grpc.ClientConnInterface, status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewDIDCommClient(cc)
	statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
		TypeId: status.Notification.ProtocolType,
		//Role:             agency.Protocol_ADDRESSEE,
		Id:               status.Notification.ProtocolId,
		NotificationTime: status.Notification.Timestamp,
	})
	err2.Check(err)
	fmt.Println("** protocol state:", statusResult.State.State)

	stateAck := agency.ProtocolState_ACK
	if !ack {
		stateAck = agency.ProtocolState_NACK
	}
	unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
		ProtocolId: &agency.ProtocolID{
			TypeId: status.Notification.ProtocolType,
			Role:   agency.Protocol_RESUME,
			Id:     status.Notification.ProtocolId,
		},
		State: stateAck,
	})
	err2.Check(err)
	glog.Infoln("result:", unpauseResult.String())
}
