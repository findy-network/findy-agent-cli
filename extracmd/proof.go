package extracmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/completionhelp"
	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var proofCmd = &cobra.Command{
	Use:   "proof",
	Short: "sends proof request to other end of the pairwise/connection",
	Long:  ``,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(cmd.BindEnvs(envs, "user"))
		return cmd.BindEnvs(proofEnvs, "send")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		err2.Check(cmdProof.Validate())
		fmt.Println(cmdProof.Name)
		if !cmd.DryRun() {
			c.SilenceUsage = true

			r, err := cmdProof.Exec(os.Stdout)
			err2.Check(err)
			result, ok := r.(*connection.Result)
			if !ok {
				panic("programming error")
			}
			if !result.Ready {
				return errors.New("proof failed or timeout")
			}
		}
		return nil
	},
}

var cmdProof = connection.ReqProofCmd{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	proofCmd.Flags().StringVar(&cmdProof.WalletName, "wallet-name", "", "edge agent's wallet name")
	err2.Check(proofCmd.Flags().SetAnnotation("wallet-name", cobra.BashCompSubdirsInDir, completionhelp.WalletLocations()))

	proofCmd.Flags().StringVar(&cmdProof.WalletKey, "wallet-key", "", "edge agent's wallet key")
	proofCmd.Flags().StringVarP(&cmdProof.Name, "conn-id", "i", "", "connection id")
	proofCmd.Flags().StringVar(&cmdProof.Attributes, "attrs", "", "proof attributes in JSON array")

	cmd.RootCmd().AddCommand(proofCmd)
}

var proofEnvs = map[string]string{
	"conn-id": "CONNECTION_ID",
}
