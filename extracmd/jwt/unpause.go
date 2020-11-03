package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-api/grpc/agency"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/grpc/client"
	"github.com/findy-network/findy-wrapper-go/dto"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "unpause command for JWT gRPC",
	Long: `
`,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if cmd.DryRun() {
			fmt.Println(dto.ToJSON(cmdData))
			return nil
		}
		c.SilenceUsage = true

		conn := client.TryOpenConn(cmdData.CaDID, cmdData.APIService, cmdData.Port)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Second)
		defer cancel()

		didComm := agency.NewDIDCommClient(conn)
		stateAck := agency.ProtocolState_ACK
		if !ACK {
			stateAck = agency.ProtocolState_NACK
		}
		unpauseResult, err := didComm.Resume(ctx, &agency.ProtocolState{
			ProtocolId: &agency.ProtocolID{
				TypeId: agency.Protocol_PROOF,
				Role:   agency.Protocol_RESUME,
				Id:     MyProtocolID,
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

	jwtCmd.AddCommand(unpauseCmd)
}