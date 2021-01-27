package jwt

import (
	"context"
	"errors"
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-grpc/agency/client"
	"github.com/findy-network/findy-wrapper-go/dto"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

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
			fmt.Println(dto.ToJSON(CmdData))
			return nil
		}
		c.SilenceUsage = true

		baseCfg := client.BuildClientConnBase("", CmdData.APIService, CmdData.Port, nil)
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		ch, err := client.Pairwise{ID: CmdData.ConnID, Conn: conn}.Ping(ctx)
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

	JwtCmd.AddCommand(pingCmd)
}
