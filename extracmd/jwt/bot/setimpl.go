package bot

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var setImplEndpDoc = `The command sets a service implementation for the cloud agent.

The Service implementation ID stands for the protocol which is used for the
communication between cloud agent and its controller. Because we are in the
middle of the transition to gRPC API we still has to use this to allow previous
API users to work.`

var setImplEndpCmd = &cobra.Command{
	Use:   "set",
	Short: "set impl id which sets it permanently as EA endpoint",
	Long:  setImplEndpDoc,
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
		conn = client.TryOpen(jwt.CmdData.CaDID, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentClient(conn)
		r, err := agent.SetImplId(ctx, &agency.SAImplementation{
			Id: implID, Persistent: persistent})
		err2.Check(err)
		fmt.Println("implementation ID set to:", r.Id)

		return nil
	},
}

var implID string
var persistent bool

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	setImplEndpCmd.Flags().StringVarP(&implID, "id", "i", "grpc", "implementation ID for us as a Edge Agent")
	setImplEndpCmd.Flags().BoolVarP(&persistent, "persistent", "p", true, "tells to write implementation ID to wallet")
	botCmd.AddCommand(setImplEndpCmd)
}
