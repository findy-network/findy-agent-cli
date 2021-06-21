package connection

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var reqProofDoc = `Requests a proof from DIDComm's other end agent.

Note! This just a simple command to test newly created credentials and not meant
to be used in production. For example, it doesn't have proper error handling,
timeouts, etc.`

var reqProofCmd = &cobra.Command{
	Use:   "reqproof",
	Short: "Request a proof and wait status",
	Long:  reqProofDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(cmd.BindEnvs(envs, ""))
		return cmd.BindEnvs(issueEnvs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: CmdData.ConnID, Conn: conn}.ReqProof(ctx, attrJSON)
		err2.Check(err)
		for status := range ch {
			fmt.Println("proof request status:", status.State, "|", status.Info)
		}
		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	reqProofCmd.Flags().StringVar(&attrJSON, "attrs", "", "attrs json")

	ConnectionCmd.AddCommand(reqProofCmd)
}
