package connection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var statusDoc = `    CONNECT = 1;
    ISSUE = 2;
    PROPOSE_PROOFING = 3;
    TRUST_PING = 4;
    BASIC_MESSAGE = 5;
`

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Return protocol status",
	Long:  statusDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			PrintCmdData()
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildConnBase(cmd.TLSPath(), cmd.ServiceAddr(), nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		didComm := agency.NewProtocolServiceClient(conn)
		statusResult, err := didComm.Status(ctx, &agency.ProtocolID{
			TypeID: agency.Protocol_Type(MyTypeID), // casting!!!
			Role:   agency.Protocol_RESUMER,
			ID:     MyProtocolID,
		})
		try.To(err)

		bStatus := try.To1(json.Marshal(statusResult.GetStatus()))

		fmt.Println("result:", string(bStatus), statusResult.State.ProtocolID.TypeID)
		return nil
	},
}

var MyTypeID int32

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	statusCmd.Flags().StringVarP(&MyProtocolID, "id", "i", "", "protocol id for continue")
	statusCmd.Flags().Int32VarP(&MyTypeID, "type", "t", 4, "4=trust ping, 1=issue, ... see usage")

	ConnectionCmd.AddCommand(statusCmd)
}
