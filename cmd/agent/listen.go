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
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		ch := conn.ListenAndRetry(ctx,
			&agency.ClientID{ID: uuid.New().String()})

	loop:
		for {
			select {
			case status, ok := <-ch:
				if !ok {
					fmt.Println("closed from server")
					break loop
				}
				fmt.Println("listen status:",
					status.Status.Notification.ProtocolType, "|",
					status.Status.Notification.TypeID, "|",
					status.Status.Notification.ProtocolID)
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

	AgentCmd.AddCommand(listenCmd)
}