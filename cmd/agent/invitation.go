package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var invitationDoc = `Commands the cloud agent to produce invitation JSON.

If conn-id is given our end of the connection used that for naming the pairwise.
if conn-id is empty CA will genereta new UUID which will be used for both ends
of the pairwise.`

var invitationCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Print connection invitation",
	Long:  invitationDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Println("JWT:", CmdData.JWT)
			fmt.Println("Server:", cmd.ServiceAddr())
			fmt.Println("Label:", ourLabel)
			fmt.Println("ConnectionID:", connID)
			if connID == "" {
				fmt.Println("autogenerated shared connection ID")
			}
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)
		r, err := agent.CreateInvitation(ctx, &agency.InvitationBase{
			ID:    connID,
			Label: ourLabel,
		})
		err2.Check(err)

		if urlFormat {
			fmt.Print(r.URL)
		} else {
			fmt.Println(r.JSON)
		}

		return nil
	},
}

var (
	connID    string
	urlFormat bool
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	invitationCmd.Flags().BoolVarP(&urlFormat,
		"url", "u", false, "if set returns URL formatted invitation")
	invitationCmd.Flags().StringVar(&ourLabel,
		"label", "", "our Aries connection Label ")
	invitationCmd.Flags().StringVarP(&connID, "conn-id", "c", "",
		"connection id (UUID) for our end, if empty autogenerated for both")

	AgentCmd.AddCommand(invitationCmd)
}
