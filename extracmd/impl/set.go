package impl

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/cmds"
	"github.com/findy-network/findy-agent/cmds/agent/sa"
	. "github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets preregistered implementation ID for an edge agent (EA)",
	Long:  ``,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "user")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer Return(&err)

		Check(cmdSet.Validate())
		if !cmd.DryRun() {
			c.SilenceUsage = true

			Try(cmdSet.Exec(os.Stdout))
		}
		return nil

	},
}

var cmdSet = sa.EAImplCmd{
	Cmd:          cmds.Cmd{},
	EAImplID:     "",
	EAServiceURL: "",
	EAServiceKey: "",
}

func init() {
	defer Catch(func(err error) {
		fmt.Println(err)
	})

	setCmd.Flags().StringVar(&cmdSet.WalletName, "wallet-name", "", "EA's wallet name")
	Check(setCmd.Flags().SetAnnotation("wallet-name", cobra.BashCompSubdirsInDir, WalletLocations()))

	setCmd.Flags().StringVar(&cmdSet.WalletKey, "wallet-key", "", "EA's wallet key")
	setCmd.Flags().StringVarP(&cmdSet.EAImplID, "impl-id", "i", "", "implementation ID")
	setCmd.Flags().StringVarP(&cmdSet.EAServiceURL, "url", "u", "", "EA's endpoint url")
	setCmd.Flags().StringVarP(&cmdSet.EAServiceKey, "key", "k", "", "EA's endpoint key")

	implCmd.AddCommand(setCmd)
}

func WalletLocations() []string {
	defer Catch(func(err error) {
		_, _ = fmt.Fprintln(os.Stderr, err)
	})

	home := String.Try(os.UserHomeDir())
	indyWallets := filepath.Join(home, ".indy_client/wallet")

	return []string{indyWallets}
}
