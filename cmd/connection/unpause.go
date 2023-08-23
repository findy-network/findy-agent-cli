package connection

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var unpauseDoc = `Resumes a protocol with given ACK/NACK.

Protocol is defined by protocol ID and protocol type: issuing or proofing`

var unpauseCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume protocol",
	Long:  unpauseDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			PrintCmdData()
			return nil
		}
		baseCfg := try.To1(cmd.BaseCfg())
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		didComm := agency.NewProtocolServiceClient(conn)
		stateAck := agency.ProtocolState_ACK
		if !ACK {
			stateAck = agency.ProtocolState_NACK
		}

		protocolTypeID := agency.Protocol_PRESENT_PROOF
		if !isProof {
			protocolTypeID = agency.Protocol_ISSUE_CREDENTIAL
		}

		unpauseResult := try.To1(didComm.Resume(ctx, &agency.ProtocolState{
			ProtocolID: &agency.ProtocolID{
				TypeID: protocolTypeID,
				Role:   agency.Protocol_RESUMER,
				ID:     MyProtocolID,
			},
			State: stateAck,
		}))

		fmt.Println("result:", unpauseResult.String())
		return nil
	},
}

var (
	MyProtocolID string
	ACK          bool
	isProof      bool
)

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	unpauseCmd.Flags().StringVarP(&MyProtocolID, "id", "i", "", "protocol id for continue")
	unpauseCmd.Flags().BoolVarP(&ACK, "ack", "a", true, "how to proceed with the protocol")
	unpauseCmd.Flags().BoolVarP(&isProof, "proof", "p", false, "are we resuming issuing or proof")

	ConnectionCmd.AddCommand(unpauseCmd)
}
