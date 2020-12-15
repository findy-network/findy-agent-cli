package bot

import (
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var botDoc = `Bot is a family of commands to manage chat bots and clients.

Most importantly you can start a chat-service, but you can as well to start
chat client. The full client needs two different windows. One for reading
incoming messages another for sending them back in a loop.`

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "bot commands",
	Long:  botDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	jwt.JwtCmd.AddCommand(botCmd)
}

var envs = map[string]string{
	"wallet-name": "WALLET_NAME",
	"wallet-key":  "WALLET_KEY",
}
