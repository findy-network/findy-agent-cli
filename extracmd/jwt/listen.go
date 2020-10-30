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
)

// userCmd represents the user command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "listen command for JWT gRPC",
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

		conn := client.TryOpenConn(cmdData.CaDID, cmdData.APIService, cmdData.Port)
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
				fmt.Println("listen status:", status.ClientId, "|", status.Notification.TypeId)
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

	jwtCmd.AddCommand(listenCmd)
}
