package agent

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

const (
	pingTimeout     = time.Second * 4
	stressTestTimer = 1 * time.Millisecond
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start to listen cloud agent until ctrl-C",
	Long:  `Starts to listen the cloud agent and prints every notification.`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)

		var msleep func(d time.Duration)
		if stressTest {
			msleep = mySleep
		}
		conn := client.TryAuthOpenWithSleep(CmdData.JWT, baseCfg, msleep)
		defer conn.Close()

		// first let's ping our agent to get proper error message for
		// JWT authentication and we are not hurry
		timeout, timeoutCancel := context.WithTimeout(
			context.Background(), pingTimeout)
		defer timeoutCancel()

		agent := agency.NewAgentServiceClient(conn)
		err2.Empty.Try(agent.Ping(timeout, &agency.PingMsg{
			ID:             1000,
			PingController: andController,
		}))

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		ch := conn.ListenStatusAndRetry(ctx,
			&agency.ClientID{ID: uuid.New().String()})

		titleOn := false
	loop:
		for {
			select {
			case status, ok := <-ch:
				if !ok {
					break loop
				}
				if !titleOn {
					titleOn = true
					fmt.Println("ProtocolID | ProtocolType | TypeID | ConnectionID")
					fmt.Println("-------------------------------------------------")
				}
				fmt.Println(
					status.Notification.ProtocolID, "|",
					status.Notification.ProtocolType, "|",
					status.Notification.TypeID, "|",
					status.Notification.ConnectionID,
				)
			case <-intCh:
				cancel()
			}
		}

		return nil
	},
}

func mySleep(d time.Duration) {
	time.Sleep(stressTestTimer)
	glog.V(3).Infof("our own short %v sleep", stressTestTimer)
}

var stressTest bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	listenCmd.Flags().BoolVarP(&stressTest, "stress", "t", false,
		"stress mode = immediate connection retry")
	AgentCmd.AddCommand(listenCmd)
}
