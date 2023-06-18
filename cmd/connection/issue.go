package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var credDefID, attrJSON string

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a credential.",
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)
		try.To(cmd.BindEnvs(envs, ""))
		return cmd.BindEnvs(issueEnvs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err, "issuing error")

		if cmd.DryRun() {
			PrintCmdData()
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch := try.To1(client.Pairwise{ID: CmdData.ConnID, Conn: conn}.Issue(ctx, credDefID, attrJSON))

		okOutput := false
		for status := range ch {
			if status.State == agency.ProtocolState_ERR {
				fmt.Fprintln(os.Stderr, "protocol error:", status.GetInfo())
				err = fmt.Errorf("protocol error: %s", status.GetInfo())
			} else {
				okOutput = true
				fmt.Println("issue status:", status.State, "|", status.Info)
			}
		}
		if !okOutput {
			err = fmt.Errorf("no response")
		}
		return err
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
