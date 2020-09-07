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
	flags.StringVar(&aCmd.APNSP12CertFile, "apns-p12-file", "", "APNS certificate p12 file")
	flags.StringVar(&aCmd.HostAddr, "host-address", "localhost", "host address")
	flags.UintVar(&aCmd.HostPort, "host-port", 8080, "host port")
	flags.UintVar(&aCmd.ServerPort, "server-port", 8080, "server port")
	flags.StringVar(&aCmd.ServiceName, "service-name", "ca-api", "service name")
	flags.StringVar(&aCmd.PoolName, "pool-name", "findy-pool", "pool name")
	flags.Uint64Var(&aCmd.PoolProtocol, "pool-protocol", 2, "pool protocol")
	flags.StringVar(&aCmd.StewardSeed, "steward-seed", "000000000000000000000000Steward1", "steward seed")
	flags.StringVar(&aCmd.PsmDb, "psm-database-file", "findy.bolt", "state machine database's filename")
	flags.BoolVar(&aCmd.ResetData, "reset-register", false, "reset handshake register")
	flags.StringVar(&aCmd.HandshakeRegister, "register-file", "findy.json", "handshake registry's filename")
	flags.StringVar(&aCmd.WalletName, "steward-wallet-name", "", "steward wallet name")
	flags.StringVar(&aCmd.WalletPwd, "steward-wallet-key", "", "steward wallet key")
	flags.StringVar(&aCmd.StewardDid, "steward-did", "", "steward DID")
	flags.StringVar(&aCmd.ServiceName2, "protocol-path", "a2a", "URL path for A2A protocols") // agency.ProtocolPath is available
	flags.StringVar(&aCmd.Salt, "salt", "", "salt")

	p := pingAgencyCmd.Flags()
	p.StringVar(&paCmd.BaseAddr, "base-address", "http://localhost:8080", "base address of agency")

	rootCmd.AddCommand(agencyCmd)
	agencyCmd.AddCommand(startAgencyCmd)
	agencyCmd.AddCommand(pingAgencyCmd)
}
