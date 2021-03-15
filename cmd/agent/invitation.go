package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var invitationCmd = &cobra.Command{
	Use:   "invitation",
	Short: "invitation command for JWT gRPC",
	Long: `Connects the cloud agent to produce invitation JSON
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Println("JWT:", CmdData.JWT)
			fmt.Println("Server:", cmd.ServiceAddr())
			fmt.Println("Label:", ourLabel)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		agent := agency.NewAgentClient(conn)
		r, err := agent.CreateInvitation(ctx, &agency.InvitationBase{
			Id:    uuid.New().String(),
			Label: ourLabel,
		})
		err2.Check(err)
		fmt.Println(r.JsonStr)

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	invitationCmd.Flags().StringVar(&ourLabel, "label", "", "our Aries connection Label ")

	AgentCmd.AddCommand(invitationCmd)
}
