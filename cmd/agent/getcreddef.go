package agent

import (
	"context"
	"fmt"

	agency "github.com/findy-network/findy-agent-api/grpc/agency/v1"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var getCredDefDoc = `Gets cred def data from the cloud agent.
Usage:
	.. get-cred-def --id YOUR_CRED_DEF_ID`

var getCredDefCmd = &cobra.Command{
	Use:   "get-cred-def",
	Short: "Gets cred def data",
	Long:  getCredDefDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Printf("CredDefID: %s\n", CredDefID)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)
		r, err := agent.GetCredDef(ctx, &agency.CredDef{
			ID: CredDefID,
		})
		err2.Check(err)
		fmt.Println(r.Data) // plain output for pipe/filter style

		return nil
	},
}

var (
	CredDefID string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	getCredDefCmd.Flags().StringVarP(&CredDefID, "id", "i", "", "credDef ID")
	getCredDefCmd.MarkFlagRequired("id")
	AgentCmd.AddCommand(getCredDefCmd)
}
