package authn

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/findy-network/findy-agent-auth/acator/authn"
	"github.com/findy-network/findy-agent-auth/acator/enclave"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var loginDoc = `Login a new WebAuthn authenticator with given arguments.
The authenticator commands are stateless which means that the caller must keep
track all of the needed data. See more information from authn parent command.`

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with our authenticator",
	Long:  loginDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		myCmd := authnCmd
		myCmd.SubCmd = c.Name()

		// we want to use our enclave here just for testing architecture
		myCmd.SecEnclave = enclave.New(myCmd.Key)

		try.To(myCmd.Validate())

		if cmd.DryRun() {
			b, _ := json.MarshalIndent(myCmd, "", "  ")
			fmt.Println(string(b))
			return nil
		}

		r := try.To1(myCmd.Exec(os.Stdout))
		if myCmd.Legacy {
			fmt.Println(r.Token)
			return nil
		}
		var result authn.Result
		try.To(json.NewDecoder(strings.NewReader(r.Token)).Decode(&result))
		fmt.Println(result.Token)

		return nil
	},
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	acatorCmd.AddCommand(loginCmd)
}
