package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agency"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

// agencyCmd represents the agency command
var agencyCmd = &cobra.Command{
	Use:   "agency",
	Short: "Parent command for starting and pinging agency",
	Long:  `Long description & example todo`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

// startAgencyCmd represents the agency start subcommand
var startAgencyCmd = &cobra.Command{
	Use:   "start",
	Short: "Command for starting agency",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		aCmd.VersionInfo = "findy-agent-cli"
		err2.Check(aCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			aCmd.PreRun()
			err2.Try(aCmd.Exec(os.Stdout))
		}
		return nil
	},
}

// pingAgencyCmd represents the agency ping subcommand
var pingAgencyCmd = &cobra.Command{
	Use:   "ping",
	Short: "Command for pinging agency",
	Long:  `Long description & example todo`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(paCmd.Validate())
		if !rootFlags.dryRun {
			cmd.SilenceUsage = true
			err2.Try(paCmd.Exec(os.Stdout))
		}
		return nil
	},
}

var aCmd = agency.Cmd{}
var paCmd = agency.PingCmd{}

func init() {
	defer err2.CatchTrace(func(err error) {
		log.Println(err)
	})

	flags := startAgencyCmd.Flags()
	flags.StringVar(&aCmd.HostAddr, "host-addr", "localhost", "host address")
	flags.UintVar(&aCmd.HostPort, "host-port", 8080, "host port")
	flags.UintVar(&aCmd.ServerPort, "server-port", 8080, "server port")
	flags.StringVar(&aCmd.ServiceName, "service-name", "ca-api", "service name")
	flags.StringVar(&aCmd.PoolTxnName, "genesis", "configs/genesis_transactions", "pool genesis file")
	flags.StringVar(&aCmd.PoolName, "pool", "findy-pool", "pool name")
	flags.Uint64Var(&aCmd.PoolProtocol, "protocol", 2, "pool protocol")
	flags.StringVar(&aCmd.StewardSeed, "seed", "000000000000000000000000Steward1", "steward seed")
	flags.StringVar(&aCmd.PsmDb, "psmdb", "findy.bolt", "state machine db's filename")
	flags.BoolVar(&aCmd.ResetData, "reset", false, "reset register")
	flags.StringVar(&aCmd.HandshakeRegister, "register", "findy.json", "handshake registry's filename")
	flags.StringVar(&aCmd.WalletName, "steward-name", "", "steward name")
	flags.StringVar(&aCmd.WalletPwd, "steward-key", "", "steward key")
	flags.StringVar(&aCmd.StewardDid, "did", "", "steward DID")
	flags.StringVar(&aCmd.ServiceName2, "a2a", "a2a", "URL path for A2A protocols") // agency.ProtocolPath is available

	err2.Check(startAgencyCmd.MarkFlagRequired("steward-name"))
	err2.Check(startAgencyCmd.MarkFlagRequired("steward-key"))

	p := pingAgencyCmd.Flags()
	p.StringVar(&paCmd.BaseAddr, "base-addr", "localhost", "base address of agency")

	rootCmd.AddCommand(agencyCmd)
	agencyCmd.AddCommand(startAgencyCmd)
	agencyCmd.AddCommand(pingAgencyCmd)
}
