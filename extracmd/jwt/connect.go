package jwt

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var connectDoc = `Connect builds a new pairwise connection to another agent. The other agent
is specified by an invitation. The invitation can be entered in three ways:

1. As a flag string (--invitation)
   $> find-agent-cli jwt connect --invitation "{inv...}"
2. As a file name including the invitation
   $> find-agent-cli jwt connect invitation.json
3. Thru the pipe when the file name is "-":
   $> echo {invitation} | find-agent-cli jwt connect -`

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect command for JWT gRPC",
	Long:  connectDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		c.SilenceUsage = true
		if len(args) > 0 {
			if args[0] == "-" {
				invitationJSON = tryReadInvitation(os.Stdin)
			} else {
				inJson := err2.File.Try(os.Open(args[0]))
				defer inJson.Close()
				invitationJSON = tryReadInvitation(inJson)
			}
		} else if invitationJSON == "" {
			fmt.Println("CMD connect {invitationJSON|-}")
			return fmt.Errorf("invitation missing")
		}

		if cmd.DryRun() {
			fmt.Println(invitationJSON)
			return nil
		}

		baseCfg := client.BuildClientConnBase("", CmdData.APIService, CmdData.Port, nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		pw := client.Pairwise{
			Conn:  conn,
			Label: ourLabel,
		}
		connID, ch, err := pw.Connection(ctx, invitationJSON)
		err2.Check(err)
		for status := range ch {
			if status.State == agency.ProtocolState_OK {
				fmt.Printf(connID)
			} else if status.State == agency.ProtocolState_ERR {
				err2.Try(fmt.Fprintln(os.Stderr, "server error:", status.Info))
			}
		}
		return nil
	},
}

var (
	invitationJSON string
	ourLabel       string
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	connectCmd.Flags().StringVar(&invitationJSON, "invitation", "", "invitation json")
	connectCmd.Flags().StringVar(&ourLabel, "label", "", "our Aries connection Label ")

	JwtCmd.AddCommand(connectCmd)
}

// readInvitation function reads invitation json, parses it & stores it to connectionCmd.Invitation pointer
func tryReadInvitation(r io.Reader) string {
	d := err2.Bytes.Try(ioutil.ReadAll(r))
	return string(d)
}
