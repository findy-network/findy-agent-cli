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

var createCredDefDoc = `The command creates a new creddef with help of the cloud agent.

A new creddef needs a Schema ID and a tag to identify a new creddef.`

var createCredDefCmd = &cobra.Command{
	Use:   "create-cred-def",
	Short: "Creates a new creddef",
	Long:  createCredDefDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Printf("schema ID: %s, tag: %s\n", schemaID, tag)
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		agent := agency.NewAgentServiceClient(conn)
		r, err := agent.CreateCredDef(ctx, &agency.CredDefCreate{
			SchemaID: schemaID, Tag: tag})
		err2.Check(err)
		fmt.Println(r.ID) // plain output for pipes

		return nil
	},
}

var tag string

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	createCredDefCmd.Flags().StringVarP(&schemaID, "id", "i", "", "schema ID of the creaddef")
	createCredDefCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag of the creddef")
	createCredDefCmd.MarkFlagRequired("id")
	createCredDefCmd.MarkFlagRequired("tag")
	AgentCmd.AddCommand(createCredDefCmd)
}
