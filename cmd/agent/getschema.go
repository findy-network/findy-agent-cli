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

var getSchemaDoc = `Gets schema data from the cloud agent.
Usage:
	.. get-schema --schema-id YOUR_SCHEMA_ID`

var getSchemaCmd = &cobra.Command{
	Use:   "get-schema",
	Short: "Gets schema data",
	Long:  getSchemaDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Printf("schemaID: %s\n", schemaID)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		agent := agency.NewAgentServiceClient(conn)
		r, err := agent.GetSchema(ctx, &agency.Schema{
			ID: schemaID,
		})
		err2.Check(err)
		// plain output for script-friendlyness of the command
		fmt.Println(r.Data)

		return nil
	},
}

var (
	schemaID string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	getSchemaCmd.Flags().StringVarP(&schemaID, "schema-id", "i", "", "schema ID")
	getSchemaCmd.MarkFlagRequired("schema-id")
	AgentCmd.AddCommand(getSchemaCmd)
}
