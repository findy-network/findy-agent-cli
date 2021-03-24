package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var createSchemaDoc = `Creates a new schema with the cloud agent.
Sample:
	.. create-schema <flags> attr1 attr2 ... attrn.`

var createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates a new schema",
	Long:  createSchemaDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if len(args) == 0 {
			return fmt.Errorf("missing attributes")
		}
		attrs := args
		if cmd.DryRun() {
			fmt.Printf("name: %s, version: %s, attributes:\n", name, version)
			fmt.Println(attrs)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentClient(conn)
		r, err := agent.CreateSchema(ctx, &agency.SchemaCreate{
			Name:    name,
			Version: version,
			Attrs:   attrs,
		})
		err2.Check(err)
		fmt.Println(r.Id) // plain output for pipes

		return nil
	},
}

var (
	name    string
	version string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	createSchemaCmd.Flags().StringVarP(&name, "name", "a", "", "schema name")
	createSchemaCmd.Flags().StringVarP(&version, "version", "v", "", "schema version")
	createSchemaCmd.MarkFlagRequired("name")
	createSchemaCmd.MarkFlagRequired("version")
	AgentCmd.AddCommand(createSchemaCmd)
}
