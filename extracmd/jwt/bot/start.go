package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/findy-network/findy-grpc/agency/client/chat"
	"github.com/findy-network/findy-grpc/agency/fsm"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var startCmdDoc = `The command starts a multi-tenant chat bot service.

The chat bot can serve what ever purpose it is programmed. The programming is
done thru state machines. The machines can be declared either YAML or JSON. The
specification for the state machine language can be found from
  [todo URL here when spec is ready]`

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a chat bot from state machine file",
	Long:  startCmdDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		var md fsm.MachineData
		if len(args) == 0 || (len(args) > 0 && args[0] == "-") {
			md, err = chat.LoadFSMMachineData(fType, os.Stdin)
			err2.Check(err)
		} else {
			fsmFile := args[0]
			f := err2.File.Try(os.Open(fsmFile))
			defer f.Close()
			md, err = chat.LoadFSMMachineData(fsmFile, f)
			err2.Check(err)
		}

		if cmd.DryRun() {
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", jwt.CmdData.APIService,
			jwt.CmdData.Port, nil)
		conn = client.TryAuthOpen(jwt.CmdData.JWT, baseCfg)
		defer conn.Close()

		// Handle graceful shutdown
		intCh := make(chan os.Signal, 1)
		signal.Notify(intCh, syscall.SIGTERM)
		signal.Notify(intCh, syscall.SIGINT)

		chat.Bot{
			Conn:        conn,
			MachineData: md,
		}.Run(intCh)

		return nil
	},
}

var (
	conn  client.Conn
	ack   bool
	fType string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	startCmd.Flags().StringVarP(&fType, "type", "t", ".yaml", "file type used for state machine load")
	botCmd.AddCommand(startCmd)
}
