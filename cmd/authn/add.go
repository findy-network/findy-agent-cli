package authn

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var addDoc = `Adds a new WebAuthn authenticator with given arguments to current account.
The account is specified with --jwt argument. Note! That the authenticator
commands are stateless which means that the caller must keep track all of the
needed data. See more information from authn parent command.`

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new authenticator to JWT specified account",
	Long:  addDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		myCmd := authnCmd
		myCmd.SubCmd = "register" // register is the right command w/ --jwt

		try.To(myCmd.Validate())
		if !cmd.DryRun() {
			r := try.To1(myCmd.Exec(os.Stdout))
			fmt.Println(r.Token)
		} else {
			b, _ := json.MarshalIndent(myCmd, "", "  ")
			fmt.Println(string(b))
		}

		return nil
	},
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	addCmd.Flags().StringVar(&authnCmd.Token, "jwt", authnCmd.Token,
		cmd.FlagInfo("Existing token to register a NEW authenticator", "", envs["jwt"]))
	addCmd.MarkFlagRequired("jwt")
	acatorCmd.AddCommand(addCmd)
}
