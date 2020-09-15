package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/findy-network/findy-agent/agent/utils"
	"github.com/findy-network/findy-agent/cmds/agent"
	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// connectCmd represents the connect subcommand
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Command for creating a2a connection between 2 agents",
	Long: `
Command for creating a2a connection between 2 agents. 

To use invitation file, pass file as command argument.
E.g. findy-agent-cli user connect path/to/invitationFile

You can also read invitation json from standard input.
E.g. findy-agent-cli user connect - {invitationJson}

Example
	findy-agent-cli user connect \
		--wallet-name MyWallet \
		--wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--name my_pairwise_name \
		--key my_pairwise_key \
		--endpoint pairwise_endpoint
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("endpoint", envPrefix+"_CONNECT_ENDPOINT"))
		err2.Check(viper.BindEnv("name", envPrefix+"_CONNECT_PAIRWISE_NAME"))
		err2.Check(viper.BindEnv("key", envPrefix+"_CONNECT_PAIRWISE_KEY"))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if len(args) > 0 {
			if args[0] == "-" {
				err2.Check(readInvitation(os.Stdin))
			} else {
				invitationFile := args[0]
				f := err2.File.Try(os.Open(invitationFile))
				defer f.Close()
				err2.Check(readInvitation(f))
			}
		} else {
			connectionCmd.Label = pwName
			connectionCmd.ID = utils.UUID()
			connectionCmd.ServiceEndpoint = pwEndp
			connectionCmd.RecipientKeys = []string{pwKey}
		}
		connectionCmd.WalletName = cFlags.WalletName
		connectionCmd.WalletKey = cFlags.WalletKey
		err2.Check(connectionCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			r, err := connectionCmd.Exec(os.Stdout)
			err2.Check(err)
			result := r.(*connection.Result)
			if result.Ready {
				fmt.Printf("connection [%s] ready\n", connectionCmd.ID)
			} else {
				fmt.Println("connection started by task id:", result.TaskID)
			}
		}
		return nil
	},
}

// readInvitation function reads invitation json, parses it & stores it to connectionCmd.Invitation pointer
func readInvitation(r io.Reader) (err error) {
	defer err2.Return(&err)
	d := err2.Bytes.Try(ioutil.ReadAll(r))
	fmt.Println(string(d))
	err2.Check(json.Unmarshal(d, &connectionCmd.Invitation))
	return nil
}

var (
	pwEndp string
	pwName string
	pwKey  string

	connectionCmd agent.ConnectionCmd
)

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	flags := connectCmd.Flags()

	flags.StringVar(&pwEndp, "endpoint", "", "pairwise endpoint, ENV variable: "+envPrefix+"_CONNECT_ENDPOINT")
	flags.StringVar(&pwName, "name", "", "name of the pairwise connection, ENV variable: "+envPrefix+"_CONNECT_PAIRWISE_NAME")
	flags.StringVar(&pwKey, "key", "", "pairwise endpoint key, ENV variable: "+envPrefix+"_CONNECT_PAIRWISE_KEY")

	userCmd.AddCommand(connectCmd)
	serviceCopy := *connectCmd
	serviceCmd.AddCommand(&serviceCopy)
}
