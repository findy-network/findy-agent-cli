package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/agent/ssi"
	"github.com/findy-network/findy-agent/cmds"
	"github.com/findy-network/findy-agent/cmds/agent/schema"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// schemaCmd represents the schema command
var schCmd = &cobra.Command{
	Use:   "schema",
	Short: "Parent command for operating with schemas",
	Long:  `Long description & example todo`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// schCreateCmd represents the schema create subcommand
var schCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Command for creating new schema",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		schAttrs = viper.GetStringSlice("schema-attrs")
		sch := &ssi.Schema{
			Name:    schName,
			Version: schVersion,
			Attrs:   schAttrs,
		}
		schemaCmd := schema.CreateCmd{
			Cmd: cmds.Cmd{
				WalletName: cFlags.WalletName,
				WalletKey:  cFlags.WalletKey,
			},
			Schema: sch,
		}
		err2.Check(schemaCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(schemaCmd.Exec(os.Stdout))
		}
		return nil
	},
}

// schReadCmd represents the schema read subcommand
var schReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Command for getting schema by id",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		schemaCmd := schema.GetCmd{
			Cmd: cmds.Cmd{
				WalletName: cFlags.WalletName,
				WalletKey:  cFlags.WalletKey,
			},
			ID: schID,
		}
		err2.Check(schemaCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(schemaCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var (
	schVersion string
	schName    string
	schAttrs   []string
	schTag     string
	schID      string
)

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	serviceCmd.AddCommand(schCmd)
	userCopy := *schCmd

	f := schCreateCmd.Flags()
	f.StringVar(&schVersion, "schema-v", "1.0", "schema version")
	f.StringVar(&schName, "schema-name", "", "schema name")
	f.StringSliceVar(&schAttrs, "schema-attrs", nil, "schema attributes")
	err2.Check(schCreateCmd.MarkFlagRequired("schema-name"))
	//schCreateCmd.MarkPersistentFlagRequired("schema-attrs") //todo: handle non-String required flags

	r := schReadCmd.Flags()
	r.StringVar(&schID, "schema-id", "", "schema id")
	err2.Check(schReadCmd.MarkFlagRequired("schema-id"))

	schCmd.AddCommand(schCreateCmd)
	schCmd.AddCommand(schReadCmd)
	readCopy := *schReadCmd

	userCopy.AddCommand(&readCopy)
	userCmd.AddCommand(&userCopy)
}
