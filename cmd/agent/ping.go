package agent

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var pingDoc = `Pings the cloud agent and optionally a controller.

Sample: .. ping -a  # ping the service agent as well`

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings cloud agent and optionally controller",
	Long:  pingDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Println("jwt:", CmdData.JWT)
			return nil
		}
		c.SilenceUsage = true

		// high security is not the point here, but help debugging with ID
		// that's why the security linter is suppressed for this
		rand.Seed(time.Now().UnixNano())
		const min = 10
		const max = 30
		// #nosec
		id := rand.Int31n(max-min+1) + min

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)
		r, err := agent.Ping(ctx, &agency.PingMsg{
			ID:             id,
			PingController: andController,
		})
		err2.Check(err)

		if id == r.ID {
			fmt.Println("ping OK")
		} else {
			fmt.Println("wrong ping ID from agency")
			fmt.Println("got:", r.ID, "has:", id)
		}
		return nil
	},
}

var andController bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	pingCmd.Flags().BoolVarP(&andController, "and-controller", "a", false,
		"ping service agent as well")
	AgentCmd.AddCommand(pingCmd)
}
