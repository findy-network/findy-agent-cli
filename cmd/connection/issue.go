package connection

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var credDefID, attrJSON string

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a credential.",
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

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: CmdData.ConnID, Conn: conn}.Issue(ctx, credDefID, attrJSON)
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

	issueCmd.Flags().StringVar(&attrJSON, "attrs", "",
		cmd.FlagInfo("attrs json", "", issueEnvs["attrs"]))
	issueCmd.Flags().StringVar(&credDefID, "cred-def-id", "",
		cmd.FlagInfo("cred def id", "", issueEnvs["cred-def-id"]))

	ConnectionCmd.AddCommand(issueCmd)
}

var issueEnvs = map[string]string{
	"attrs":       "ATTRS",
	"cred-def-id": "CRED_DEF_ID",
}
