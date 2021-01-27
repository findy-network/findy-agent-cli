package bot

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var chatDoc = `Starts the chat client which reads standard input and send each line
as Aries basic_message thru the pairwise connection.`

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat client to send basic messages",
	Long:  chatDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", jwt.CmdData.APIService, jwt.CmdData.Port, nil)
		conn = client.TryAuthOpen(jwt.CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			r, err := client.Pairwise{
				ID:   jwt.CmdData.ConnID,
				Conn: conn,
			}.BasicMessage(ctx, scanner.Text())
			err2.Check(err)

			for status := range r {
				fmt.Println("message status:", status.State, "|", status.Info)
			}
		}
		err2.Check(scanner.Err())

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	chatCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")
	botCmd.AddCommand(chatCmd)
}
