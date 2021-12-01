package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/findy-network/findy-common-go/std/didexchange/invitation"
	"github.com/golang/glog"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var connectDoc = `Builds a new DIDComm connection to another agent. The other agent
is specified by an invitation. The invitation can be entered in three ways:

1. As a flag string (--invitation)
   $> find-agent-cli agent connect --invitation "{inv...}"
2. As a file name including the invitation
   $> find-agent-cli agent connect invitation.json
3. Thru the pipe when the file name is "-":
   $> echo {invitation} | find-agent-cli agent connect -`

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to agent",
	Long:  connectDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Annotate("connect cmd", &err)

		c.SilenceUsage = true
		if len(args) > 0 {
			if args[0] == "-" {
				invitationStr = tryReadInvitation(os.Stdin)
			} else {
				inJSON := err2.File.Try(os.Open(args[0]))
				defer inJSON.Close()
				invitationStr = tryReadInvitation(inJSON)
			}
		} else if invitationStr == "" {
			fmt.Fprintln(os.Stderr,
				"Usage: findy-agent-cli agent connect {invitationJSON|-}")
			return fmt.Errorf("invitation missing")
		}

		glog.V(1).Infof("'%s'", invitationStr)
		invitationStr = strings.TrimSuffix(invitationStr, "\n")
		glog.V(1).Infof("'%s'", invitationStr)

		if cmd.DryRun() {
			tryPrintInvitation(invitationStr)
			return nil
		}

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn := client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		pw := client.Pairwise{
			Conn:  conn,
			Label: ourLabel,
		}
		connID, ch, err := pw.Connection(ctx, invitationStr)
		err2.Check(err)

		for status := range ch {
			if status.State == agency.ProtocolState_OK {
				fmt.Println(connID)
			} else if status.State == agency.ProtocolState_ERR {
				err2.Try(fmt.Fprintln(os.Stderr, "server error:", status.Info))
			}
		}
		return nil
	},
}

var (
	invitationStr string
	ourLabel      string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	connectCmd.Flags().StringVar(&invitationStr, "invitation", "", "invitation json")
	connectCmd.Flags().StringVar(&ourLabel, "label", "", "our Aries connection Label ")

	AgentCmd.AddCommand(connectCmd)
}

// readInvitation function reads invitation json, parses it & stores it to connectionCmd.Invitation pointer
func tryReadInvitation(r io.Reader) string {
	d := err2.Bytes.Try(ioutil.ReadAll(r))
	return string(d)
}

func tryPrintInvitation(s string) {
	var inv invitation.Invitation

	s = strings.TrimSuffix(s, "\n")
	inv, err := invitation.Translate(s)
	err2.Check(err)

	b := err2.Bytes.Try(json.MarshalIndent(inv, "", " "))
	fmt.Println(string(b))
}
