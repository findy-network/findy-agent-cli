package agent

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var createSchemaDoc = `Creates a new schema with the cloud agent.
Sample:
	.. create-schema <flags> attr1 attr2 ... attrn.`

var createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates a new schema",
	Long:  createSchemaDoc,
	Args:  cobra.MinimumNArgs(1),
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(_ *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		attrs := args
		if cmd.DryRun() {
			fmt.Printf("name: %s, version: %s, attributes:\n", name, version)
			fmt.Println(attrs)
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentServiceClient(conn)
		r := try.To1(agent.CreateSchema(ctx, &agency.SchemaCreate{
			Name:       name,
			Version:    version,
			Attributes: attrs,
		}))
		fmt.Println(r.ID) // plain output for pipes

		return nil
	},
}

var (
	name    string
	version string
)

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))
	createSchemaCmd.Flags().StringVarP(&name, "name", "a", "", "schema name")
	createSchemaCmd.Flags().StringVar(&version, "version", "", "schema version")
	createSchemaCmd.MarkFlagRequired("name")
	createSchemaCmd.MarkFlagRequired("version")
	AgentCmd.AddCommand(createSchemaCmd)
}
