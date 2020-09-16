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
	Long: `
Parent command for starting and pinging agency
	`,
	Run: func(cmd *cobra.Command, args []string) {
		SubCmdNeeded(cmd)
	},
}

var agencyStartEnvs = map[string]string{
	"apns-p12-file":       "APNS_P12_FILE",
	"host-address":        "HOST_ADDRESS",
	"host-port":           "HOST_PORT",
	"server-port":         "SERVER_PORT",
	"service-name":        "SERVICE_NAME",
	"pool-name":           "POOL_NAME",
	"pool-protocol":       "POOL_PROTOCOL",
	"steward-seed":        "STEWARD_SEED",
	"psm-database-file":   "PSM_DATABASE_FILE",
	"reset-register":      "RESET_REGISTER",
	"register-file":       "REGISTER_FILE",
	"steward-wallet-name": "STEWARD_WALLET_NAME",
	"steward-wallet-key":  "STEWARD_WALLET_KEY",
	"steward-did":         "STEWARD_DID",
	"protocol-path":       "PROTOCOL_PATH",
	"salt":                "SALT",
}

// startAgencyCmd represents the agency start subcommand
var startAgencyCmd = &cobra.Command{
	Use:   "start",
	Short: "Command for starting agency",
	Long: `
Start command for findy agency server.

Example
	findy-agent-cli agency start \
		--pool-name findy \
		--steward-wallet-name sovrin_steward_wallet \
		--steward-wallet-key 6cih1cVgRH8...dv67o8QbufxaTHot3Qxp \
		--steward-did Th7MpTaRZVRYnPiabds81Y \
		--salt mySalt
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return bindEnvs(agencyStartEnvs, "AGENCY")
	},
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

var agencyPingEnvs = map[string]string{
	"base-address": "PING_BASE_ADDRESS",
}

// pingAgencyCmd represents the agency ping subcommand
var pingAgencyCmd = &cobra.Command{
	Use:   "ping",
	Short: "Command for pinging agency",
	Long: `
Pings agency.
If agency works fine, ping ok with server's host address is printed.

Example
	findy-agent-cli agency ping \
		--base-address http://localhost:8080
	`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return bindEnvs(agencyPingEnvs, "AGENCY")
	},
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
	flags.StringVar(&aCmd.APNSP12CertFile, "apns-p12-file", "", flagInfo("APNS certificate p12 file", agencyCmd.Name(), agencyStartEnvs["apns-p12-file"]))
	flags.StringVar(&aCmd.HostAddr, "host-address", "localhost", flagInfo("host address", agencyCmd.Name(), agencyStartEnvs["host-address"]))
	flags.UintVar(&aCmd.HostPort, "host-port", 8080, flagInfo("host port", agencyCmd.Name(), agencyStartEnvs["host-port"]))
	flags.UintVar(&aCmd.ServerPort, "server-port", 8080, flagInfo("server port", agencyCmd.Name(), agencyStartEnvs["server-port"]))
	flags.StringVar(&aCmd.ServiceName, "service-name", "ca-api", flagInfo("service name", agencyCmd.Name(), agencyStartEnvs["service-name"]))
	flags.StringVar(&aCmd.PoolName, "pool-name", "findy-pool", flagInfo("pool name", agencyCmd.Name(), agencyStartEnvs["pool-name"]))
	flags.Uint64Var(&aCmd.PoolProtocol, "pool-protocol", 2, flagInfo("pool protocol", agencyCmd.Name(), agencyStartEnvs["pool-protocol"]))
	flags.StringVar(&aCmd.StewardSeed, "steward-seed", "000000000000000000000000Steward1", flagInfo("steward seed", agencyCmd.Name(), agencyStartEnvs["steward-seed"]))
	flags.StringVar(&aCmd.PsmDb, "psm-database-file", "findy.bolt", flagInfo("state machine database's filename", agencyCmd.Name(), agencyStartEnvs["psm-database-file"]))
	flags.BoolVar(&aCmd.ResetData, "reset-register", false, flagInfo("reset handshake register", agencyCmd.Name(), agencyStartEnvs["reset-register"]))
	flags.StringVar(&aCmd.HandshakeRegister, "register-file", "findy.json", flagInfo("handshake registry's filename", agencyCmd.Name(), agencyStartEnvs["register-file"]))
	flags.StringVar(&aCmd.WalletName, "steward-wallet-name", "", flagInfo("steward wallet name", agencyCmd.Name(), agencyStartEnvs["steward-wallet-name"]))
	flags.StringVar(&aCmd.WalletPwd, "steward-wallet-key", "", flagInfo("steward wallet key", agencyCmd.Name(), agencyStartEnvs["steward-wallet-key"]))
	flags.StringVar(&aCmd.StewardDid, "steward-did", "", flagInfo("steward DID", agencyCmd.Name(), agencyStartEnvs["steward-did"]))
	flags.StringVar(&aCmd.ServiceName2, "protocol-path", "a2a", flagInfo("URL path for A2A protocols", agencyCmd.Name(), agencyStartEnvs["protocol-path"])) // agency.ProtocolPath is available
	flags.StringVar(&aCmd.Salt, "salt", "", flagInfo("salt", agencyCmd.Name(), agencyStartEnvs["salt"]))

	p := pingAgencyCmd.Flags()
	p.StringVar(&paCmd.BaseAddr, "base-address", "http://localhost:8080", flagInfo("base address of agency", agencyCmd.Name(), agencyPingEnvs["base-address"]))

	rootCmd.AddCommand(agencyCmd)
	agencyCmd.AddCommand(startAgencyCmd)
	agencyCmd.AddCommand(pingAgencyCmd)
}
