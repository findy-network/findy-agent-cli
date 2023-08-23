package bot

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/golang/glog"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var chatDoc = `Starts the chat client which reads standard input and send each line
as Aries basic_message thru the pairwise connection.`

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat client to send basic messages",
	Long:  chatDoc,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			PrintCmdData()
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // for server side stops, for proper cleanup

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			r := try.To1(client.Pairwise{
				ID:   CmdData.ConnID,
				Conn: conn,
			}.BasicMessage(ctx, scanner.Text()))

			for status := range r {
				glog.V(2).Infoln("message status:", status.State, "|", status.Info)
			}
		}
		try.To(scanner.Err())

		return nil
	},
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))
	chatCmd.Flags().BoolVarP(&ack, "reply_ack", "a", true, "used reply ack for all request")
	botCmd.AddCommand(chatCmd)
}
