package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/findy-network/findy-common-go/agency/client/chat"
	"github.com/findy-network/findy-common-go/agency/fsm"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var startCmdDoc = `The command starts a multi-tenant chat bot service.

	findy-agent-cli bot start <filename.yaml/json|->

If - is given instead of the FSM declaration file, it will be read from stdio.

The chat bot can serve what ever purpose it is programmed. The programming is
done thru state machines. The machines can be declared either YAML or JSON. The
specification for the state machine language can be found from:
  [todo URL here when spec is ready]`

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a chat bot from state machine file",
	Long:  startCmdDoc,
	Args:  cobra.ExactArgs(1),
	RunE:  start,
}

func start(_ *cobra.Command, args []string) (err error) {
	defer err2.Handle(&err)

	var md fsm.MachineData
	if len(args) == 0 || (len(args) > 0 && args[0] == "-") {
		md = try.To1(chat.LoadFSMMachineData(fType, os.Stdin))
	} else {
		md = *try.To1(loadFSM(args[0]))
	}

	var mdService *fsm.MachineData
	if serviceFSM != "" {
		mdService = try.To1(loadFSM(serviceFSM))
	}

	if cmd.DryRun() {
		PrintCmdData()
		return nil
	}
	baseCfg := try.To1(cmd.BaseCfg())
	conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
	defer conn.Close()

	// Handle graceful shutdown
	intCh := make(chan os.Signal, 1)
	signal.Notify(intCh, syscall.SIGTERM, syscall.SIGINT)

	chat.Bot{
		Conn:        conn,
		MachineData: md,
		ServiceFSM:  mdService,
	}.Run(intCh)

	return nil
}

var (
	conn       client.Conn
	fType      string
	serviceFSM string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	startCmd.Flags().StringVarP(&fType, "type", "t", ".yaml", "file type used for state machine load")
	startCmd.Flags().StringVar(&serviceFSM, "service-fsm", "", "FSM file for service level state machine load")
	botCmd.AddCommand(startCmd)
}

func loadFSM(fsmFile string) (md *fsm.MachineData, err error) {
	defer err2.Handle(&err)

	f := try.To1(os.Open(fsmFile))
	defer f.Close()
	m := try.To1(chat.LoadFSMMachineData(fsmFile, f))
	return &m, nil
}
