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

var statusDoc = `    CONNECT = 0;
    ISSUE = 1;
    PROPOSE_PROOFING = 2;
    TRUST_PING = 3;
    BASIC_MESSAGE = 4;
`

// userCmd represents the user command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status command for JWT gRPC",
	Long:  statusDoc,
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
		statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
			TypeId: agency.Protocol_Type(MyTypeID),
			Role:   agency.Protocol_RESUME,
			Id:     MyProtocolID,
		})
		err2.Check(err)

		fmt.Println("result:", statusResult.Message)
		return nil
	},
}

var MyTypeID int32

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	statusCmd.Flags().StringVarP(&MyProtocolID, "id", "i", "", "protocol id for continue")
	statusCmd.Flags().Int32VarP(&MyTypeID, "type", "t", 2, "3 req proof, 1 issue, see usage")

	jwtCmd.AddCommand(statusCmd)
}