package connection

import (
	"context"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/findy-network/findy-common-go/dto"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume protocol",
	Long:  `Resumes a protocol with given ACK/NACK.`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Println(dto.ToJSON(CmdData))
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase("", cmd.ServiceAddr(), nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		didComm := agency.NewProtocolServiceClient(conn)
		stateAck := agency.ProtocolState_ACK
		if !ACK {
			stateAck = agency.ProtocolState_NACK
		}
		unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
			ProtocolID: &agency.ProtocolID{
				TypeID: agency.Protocol_PRESENT_PROOF,
				Role:   agency.Protocol_RESUMER,
				ID:     MyProtocolID,
			},
			State: stateAck,
		})
		err2.Check(err)

		fmt.Println("result:", unpauseResult.String())
		return nil
	},
}

var (
	MyProtocolID string
	ACK          bool
)

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	unpauseCmd.Flags().StringVarP(&MyProtocolID, "id", "i", "", "protocol id for continue")
	unpauseCmd.Flags().BoolVarP(&ACK, "ack", "a", true, "how to proceed with the protocol")

	ConnectionCmd.AddCommand(unpauseCmd)
}
