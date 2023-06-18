package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

// #nosec G101
var getCredDefDoc = `Gets cred def data from the cloud agent.
Usage:
	.. get-cred-def --id YOUR_CRED_DEF_ID`

var getCredDefCmd = &cobra.Command{
	Use:   "get-cred-def",
	Short: "Gets cred def data",
	Long:  getCredDefDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)
		try.To(cmd.BindEnvs(envs, ""))
		return cmd.BindEnvs(credDefEnvs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			fmt.Printf("CredDefID: %s\n", CredDefID)
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)

		for left := pollTimeout - wait; left >= 0; left -= wait {
			r, err := agent.GetCredDef(ctx, &agency.CredDef{
				ID: CredDefID,
			})

			if err == nil {
				// plain output for script-friendlyness of the command
				fmt.Println(r.Data)
				return nil
			}

			// if wait time is 0 we don't poll, but run once
			if wait == 0 {
				return err
			}

			time.Sleep(wait)
		}

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

	flags := getCredDefCmd.Flags()
	flags.StringVarP(&CredDefID, "id", "i", "",
		cmd.FlagInfo("credDef ID", "", credDefEnvs["id"]))

	flags.DurationVarP(&wait, "wait", "w", time.Second, "sleep between polls, 0 == no poll")
	flags.DurationVar(&pollTimeout, "timeout", 10*time.Second, "how long to poll until give up")

	getCredDefCmd.MarkFlagRequired("id")

	AgentCmd.AddCommand(getCredDefCmd)
}

var credDefEnvs = map[string]string{
	"id": "CRED_DEF_ID",
}
