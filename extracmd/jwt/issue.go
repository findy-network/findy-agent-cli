package jwt

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/grpc/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var credDefID, attrJSON string

// userCmd represents the user command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "issue credential command for JWT gRPC",
	Long: `
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(cmd.BindEnvs(envs, ""))
		return cmd.BindEnvs(issueEnvs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Annotate("issuing error", &err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		addr := fmt.Sprintf("%s:%d", cmdData.APIService, cmdData.Port)
		conn, err := client.NewClient(cmdData.CaDID, addr)
		err2.Check(err)

		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: cmdData.ConnID}.Issue(ctx, credDefID, attrJSON)
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

	issueCmd.Flags().StringVar(&attrJSON, "attrs", "", "attrs json")
	issueCmd.Flags().StringVar(&credDefID, "cred-def-id", "", "cred def id")

	jwtCmd.AddCommand(issueCmd)
}

var issueEnvs = map[string]string{
	"attrs":       "ATTRS",
	"cred-def-id": "CRED_DEF_ID",
}
