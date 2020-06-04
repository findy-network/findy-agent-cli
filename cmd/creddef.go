package cmd

import (
	"log"
	"os"

	"github.com/lainio/err2"
	"github.com/optechlab/findy-agent/cmds"
	"github.com/optechlab/findy-agent/cmds/agent/creddef"
	"github.com/spf13/cobra"
)

// creddefCmd represents the creddef command
var creddefCmd = &cobra.Command{
	Use:   "creddef",
	Short: "Parent command for operating with Credential definations",
	Long:  `Long description & example todo`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// createCreddefCmd represents the creddef create subcommand
var createCreddefCmd = &cobra.Command{
	Use:   "create",
	Short: "Command for creating new credential definition",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		credDefCmd := creddef.CreateCmd{
			Cmd: cmds.Cmd{
				WalletName: cFlags.WalletName,
				WalletKey:  cFlags.WalletKey,
			},
			SchemaID: schID,
			Tag:      credDefTag,
		}
		err2.Check(credDefCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(credDefCmd.Exec(os.Stdout))
		}
		return nil
	},
}

// readCreddefCmd represents the creddef read subcommand
var readCreddefCmd = &cobra.Command{
	Use:   "read",
	Short: "Command for getting credential definition by id",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		credDefCmd := creddef.GetCmd{
			Cmd: cmds.Cmd{
				WalletName: cFlags.WalletName,
				WalletKey:  cFlags.WalletKey,
			},
			ID: credDefID,
		}
		err2.Check(credDefCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(credDefCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var (
	credDefTag string
	credDefID  string
)

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	serviceCmd.AddCommand(creddefCmd)
	userCopy := *creddefCmd

	f := creddefCmd.PersistentFlags()
	f.StringVar(&schID, "schema-id", "", "schema ID")
	err2.Check(creddefCmd.MarkPersistentFlagRequired("schema-id"))

	c := createCreddefCmd.Flags()
	c.StringVar(&credDefTag, "tag", "", "cred def tag")
	err2.Check(createCreddefCmd.MarkFlagRequired("tag"))

	r := readCreddefCmd.Flags()
	r.StringVar(&credDefID, "creddef-id", "", "cred def id")
	err2.Check(readCreddefCmd.MarkFlagRequired("creddef-id"))

	creddefCmd.AddCommand(readCreddefCmd)
	readCopy := *readCreddefCmd
	creddefCmd.AddCommand(createCreddefCmd)

	userCopy.AddCommand(&readCopy)
	userCmd.AddCommand(&userCopy)
}
