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

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat client",
	Long: `
`,
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
		conn = client.TryOpen(jwt.CmdData.CaDID, baseCfg)
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
