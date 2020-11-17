package jwt

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var reqProofCmd = &cobra.Command{
	Use:   "reqproof",
	Short: "request proof command",
	Long: `
`,
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

		baseCfg := client.BuildClientConnBase("", cmdData.APIService, cmdData.Port, nil)
		conn = client.TryOpen(cmdData.CaDID, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: cmdData.ConnID, Conn: conn}.ReqProof(ctx, attrJSON)
		err2.Check(err)
		for status := range ch {
			fmt.Println("issue status:", status.State, "|", status.Info)
		}
		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	reqProofCmd.Flags().StringVar(&attrJSON, "attrs", "", "attrs json")

	jwtCmd.AddCommand(reqProofCmd)
}
