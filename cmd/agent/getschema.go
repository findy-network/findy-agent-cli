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

var getSchemaDoc = `Gets schema data from the cloud agent.
Usage:
	.. get-schema --schema-id YOUR_SCHEMA_ID`

var getSchemaCmd = &cobra.Command{
	Use:   "get-schema",
	Short: "Gets schema data",
	Long:  getSchemaDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		try.To(cmd.BindEnvs(envs, ""))
		return cmd.BindEnvs(getSchemaEnvs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Printf("schemaID: %s\n", schemaID)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)

		for left := pollTimeout - wait; left >= 0; left -= wait {
			r, err := agent.GetSchema(ctx, &agency.Schema{
				ID: schemaID,
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
	schemaID string

	wait        time.Duration
	pollTimeout time.Duration
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	flags := getSchemaCmd.Flags()
	flags.StringVarP(&schemaID, "schema-id", "i", "",
		cmd.FlagInfo("schema ID", "", getSchemaEnvs["schema-id"]))

	flags.DurationVarP(&wait, "wait", "w", time.Second, "sleep between polls, 0 == no poll")
	flags.DurationVar(&pollTimeout, "timeout", 10*time.Second, "how long to poll until give up")

	getSchemaCmd.MarkFlagRequired("schema-id")
	AgentCmd.AddCommand(getSchemaCmd)
}

var getSchemaEnvs = map[string]string{
	"schema-id": "SCHEMA_ID",
	"id":        "SCHEMA_ID",
}
