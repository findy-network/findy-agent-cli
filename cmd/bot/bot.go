package bot

import (
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var botDoc = `Bot is a family of commands to manage chat bots and clients.

Most importantly you can start a chat-service, but you can as well to start
chat client. The full client needs two different windows. One for reading
incoming messages another for sending them back in a loop.`

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Manage Bot",
	Long:  botDoc,
	PreRunE: func(*cobra.Command, []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	Run: func(c *cobra.Command, _ []string) {
		cmd.SubCmdNeeded(c)
	},
}

var CmdData = struct {
	ConnID string
	JWT    string
}{}

func PrintCmdData() {
	cb := try.To1(yaml.Marshal(CmdData))
	fmt.Println(string(cb))
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	flags := botCmd.PersistentFlags()
	flags.StringVar(&CmdData.JWT, "jwt", "",
		cmd.FlagInfo("agent JWT token", "", envs["jwt"]))
	flags.StringVar(&CmdData.ConnID, "conn-id", "",
		cmd.FlagInfo("connection id aka pairwise id", "", envs["conn-id"]))

	cmd.RootCmd().AddCommand(botCmd)
}

var envs = map[string]string{
	"jwt":     "JWT",
	"conn-id": "CONN_ID",
}
