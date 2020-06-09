package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/pool"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Parent command for pool commands",
	Long:  `Long description & example todo`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// createPoolCmd represents the pool create subcommand
var createPoolCmd = &cobra.Command{
	Use:   "create",
	Short: "Command for creating creating pool",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		Cmd := pool.CreateCmd{
			Name: poolName,
			Txn:  poolGen,
		}
		err2.Check(Cmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(Cmd.Exec(os.Stdout))
		}
		return nil
	},
}

// pingPoolCmd represents the pool ping subcommand
var pingPoolCmd = &cobra.Command{
	Use:   "ping",
	Short: "Command for pinging pool",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		Cmd := pool.PingCmd{
			Name: poolName,
		}
		err2.Check(Cmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(Cmd.Exec(os.Stdout))
		}
		return nil
	},
}

var (
	poolName string
	poolGen  string
)

func init() {
	defer err2.Catch(func(err error) {
		log.Println(err)
	})

	f := poolCmd.PersistentFlags()
	f.StringVar(&poolName, "pool-name", "", "name of the pool")
	err2.Check(poolCmd.MarkPersistentFlagRequired("pool-name"))

	c := createPoolCmd.Flags()
	c.StringVar(&poolGen, "pool-genesis", "", "pool genesis file")
	err2.Check(createPoolCmd.MarkFlagRequired("pool-genesis"))

	rootCmd.AddCommand(poolCmd)
	poolCmd.AddCommand(createPoolCmd)
	poolCmd.AddCommand(pingPoolCmd)
}
