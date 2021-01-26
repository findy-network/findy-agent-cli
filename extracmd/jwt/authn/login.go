package authn

import (
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var loginDoc = `Starts the chat client which reads standard input and send each line
as Aries basic_message thru the pairwise connection.`

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with our authenticator",
	Long:  loginDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "authn")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		myCmd := authnCmd
		myCmd.SubCmd = c.Name()

		err2.Check(myCmd.Validate())
		if !cmd.DryRun() {
			r, err := myCmd.Exec(os.Stdout)
			err2.Check(err)
			fmt.Println(r.Token)
		}

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	acatorCmd.AddCommand(loginCmd)
}
