package bot

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/findy-network/findy-grpc/agency/client/chat"
	"github.com/findy-network/findy-grpc/agency/fsm"
	"github.com/golang/glog"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startCmdDoc = `bot cmd starts a chat bot.`
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a chat bot",
	Long:  startCmdDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		var m *fsm.Machine
		if len(args) > 0 {
			if args[0] == "-" {
				m, err = chat.LoadFSM(fType, os.Stdin)
				err2.Check(err)
			} else {
				fsmFile := args[0]
				f := err2.File.Try(os.Open(fsmFile))
				defer f.Close()
				m, err = chat.LoadFSM(fsmFile, f)
				err2.Check(err)
			}
		}
		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", jwt.CmdData.APIService,
			jwt.CmdData.Port, nil)
		conn = client.TryOpen(jwt.CmdData.CaDID, baseCfg)
		defer conn.Close()

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		chat.Bot{Conn: conn, Machine: m}.Run(intCh)

		return nil
	},
}

var conn client.Conn
var ack bool
var fType string

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	startCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")
	startCmd.Flags().StringVarP(&fType, "type", "t", ".yaml", "file type used for state machine load")
	botCmd.AddCommand(startCmd)
}

func _(conn client.Conn, status *agency.AgentStatus, _ bool) {
	if status.Notification.ProtocolType == agency.Protocol_CONNECT {

	} else if status.Notification.ProtocolType == agency.Protocol_BASIC_MESSAGE {
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
		ch, err := client.Pairwise{
			ID:   status.Notification.ConnectionId,
			Conn: conn,
		}.BasicMessage(context.Background(), statusResult.GetBasicMessage().Content)
		err2.Check(err)
		for state := range ch {
			glog.V(1).Infoln("BM send state:", state.State, "|", state.Info)
		}
	}
}

func reply(conn *grpc.ClientConn, status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	c := agency.NewAgentClient(conn)
	cid, err := c.Give(ctx, &agency.Answer{
		Id:       status.Notification.Id,
		ClientId: status.ClientId,
		Ack:      ack,
		Info:     "testing says hello!",
	})
	err2.Check(err)
	glog.Infof("Sending the answer (%s) send to client:%s\n", status.Notification.Id, cid.Id)
}

func resume(conn *grpc.ClientConn, status *agency.AgentStatus, ack bool) {
	ctx := context.Background()
	didComm := agency.NewDIDCommClient(conn)
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
