package jwt

import (
	"context"
	"errors"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent/grpc/client"
	"github.com/findy-network/findy-wrapper-go/dto"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping command for JWT gRPC",
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

		baseCfg := client.BuildClientConnBase("", cmdData.APIService, cmdData.Port, nil)
		conn = client.TryOpen(cmdData.CaDID, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: cmdData.ConnID, Conn:conn}.Ping(ctx)
		err2.Check(err)
		for status := range ch {
			fmt.Println("ping status:", status.State, "|", status.Info)
			if !client.OkStatus(status) {
				panic(errors.New("error in panic"))
			}
		}
		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	jwtCmd.AddCommand(pingCmd)
}
