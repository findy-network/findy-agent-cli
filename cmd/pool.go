package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/pool"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// poolCmd represents the pool command
var poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "Parent command for pool commands",
	Long: `
Parent command for pool commands
	`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// createPoolCmd represents the pool create subcommand
var createPoolCmd = &cobra.Command{
	Use:   "create",
	Short: "Command for creating creating pool",
	Long: `
Command for creating creating pool

Example
	findy-agent-cli ledger pool create \
		--name findy-pool \
		--genesis-txn-file my-genesis-txn-file
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("name", envPrefix+"_POOL_NAME"))
		err2.Check(viper.BindEnv("genesis-txn-file", envPrefix+"_POOL_GENESIS_TXN_FILE"))
		return nil
	},
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
	Long: `
Command for pinging pool

Example
	findy-agent-cli ledger pool ping \
		--name findy-pool
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("name", envPrefix+"_POOL_NAME"))
		return nil
	},
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
	f.StringVar(&poolName, "name", "", "name of the pool, ENV variable: "+envPrefix+"_POOL_NAME")

	c := createPoolCmd.Flags()
	c.StringVar(&poolGen, "genesis-txn-file", "", "pool genesis transactions file, ENV variable: "+envPrefix+"_POOL_GENESIS_TXN_FILE")

	ledgerCmd.AddCommand(poolCmd)
	poolCmd.AddCommand(createPoolCmd)
	poolCmd.AddCommand(pingPoolCmd)
}
