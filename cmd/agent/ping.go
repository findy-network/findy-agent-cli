package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/findy-network/findy-common-go/x"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var pingDoc = `Pings the cloud agent and optionally a controller.

Sample: .. ping -a  # ping the service agent as well`

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings cloud agent and optionally controller",
	Long:  pingDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		if cmd.DryRun() {
			fmt.Println("jwt:", CmdData.JWT)
			return nil
		}

		baseCfg := try.To1(cmd.BaseCfg())
		startTime = time.Now()
		conn = client.TryAuthOpen(CmdData.JWT, baseCfg)
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		info = "UNKNOWN"
		if count != 0 {
			timedPing(ctx)
		} else {
			agent := agency.NewAgentServiceClient(conn)
			r := try.To1(agent.Enter(ctx, &agency.ModeCmd{
				TypeID: agency.ModeCmd_NONE,
			}))
			info = r.Info
		}
		fmt.Println("Agent registered by name:", info)
		return nil
	},
}

var (
	andController bool
	count         int
	preCount      int
	elapsedTotal  time.Duration
	startTime     time.Time
	info          string
)

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))
	pingCmd.Flags().BoolVarP(&andController, "and-controller", "a", false,
		"ping service agent as well")
	pingCmd.Flags().IntVarP(&count, "count", "c", 0,
		"how many times we should ping, 0 ping once without timing")
	AgentCmd.AddCommand(pingCmd)
}

func timedPing(ctx context.Context) {
	if count == 0 {
		return
	}
	agent := agency.NewAgentServiceClient(conn)
	elapsedTotal = time.Since(startTime)
	fmt.Printf("Open connection time %v\n", elapsedTotal)

	printProgressBar()
	var (
		maxElabsed time.Duration
		maxIndex   = 0
	)
	for i := 0; i < count; i++ {
		startTime = time.Now()
		r := try.To1(agent.Enter(ctx, &agency.ModeCmd{
			TypeID: agency.ModeCmd_NONE,
		}))
		elapsed := time.Since(startTime)
		if elapsed > maxElabsed {
			maxElabsed = elapsed
			maxIndex = i + 1
		}
		elapsedTotal += elapsed
		if i == 0 {
			info = r.Info
		}
		printProgress(i)
	}
	fmt.Printf("\nMax ping time: %v at ping #%d", maxElabsed, maxIndex)
	if count > 1 {
		fmt.Printf("\nNormalized meantime %v",
			(elapsedTotal-maxElabsed)/time.Duration(count-1))
	}
	fmt.Printf("\nMeantime %v\n\n", elapsedTotal/time.Duration(count))
}

func printProgressBar() {
	if count == 0 {
		return
	}
	fmt.Println("pinging", count, "times")
	countNormalized := x.Whom(count > 100, 100, count)
	for i := 0; i < countNormalized; i++ {
		fmt.Print(".")
	}
	fmt.Print("\r")
}

func printProgress(i int) {
	if count == 0 {
		return
	}
	if count > 100 {
		normI := 100 * i / count
		if normI != preCount || i == count-1 {
			fmt.Print("o")
			preCount = normI
		}
	} else {
		fmt.Print("o")
	}
}
