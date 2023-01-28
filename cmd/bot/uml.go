package bot

import (
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client/chat"
	"github.com/findy-network/findy-common-go/agency/fsm"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var umlCmdDoc = `The command umls a multi-tenant chat bot service.

The chat bot can serve what ever purpose it is programmed. The programming is
done thru state machines. The machines can be declared either YAML or JSON. The
specification for the state machine language can be found from
  [todo URL here when spec is ready]`

var umlCmd = &cobra.Command{
	Use:   "uml",
	Short: "uml a chat bot from state machine file",
	Long:  umlCmdDoc,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		var md fsm.MachineData
		if len(args) == 0 || (len(args) > 0 && args[0] == "-") {
			md = try.To1(chat.LoadFSMMachineData(fType, os.Stdin))
		} else {
			fsmFile := args[0]
			f := try.To1(os.Open(fsmFile))
			defer f.Close()
			md = try.To1(chat.LoadFSMMachineData(fsmFile, f))
		}

		if cmd.DryRun() {
			PrintCmdData()
			return nil
		}
		c.SilenceUsage = true

		m := fsm.NewMachine(md)
		url := try.To1(fsm.GenerateURL("svg", m))
		fmt.Println(url)

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})
	umlCmd.Flags().StringVarP(&fType, "type", "t", ".yaml",
		"file type used for state machine load")
	botCmd.AddCommand(umlCmd)
}
