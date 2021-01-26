package authn

import (
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var registerDoc = `Starts the chat client which reads standard input and send each line
as Aries basic_message thru the pairwise connection.`

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "registers our authenticator",
	Long:  registerDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		myCmd := authnCmd
		myCmd.SubCmd = c.Use

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

	acatorCmd.AddCommand(registerCmd)
}
