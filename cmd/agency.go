package cmd

import (
	"log"
	"os"

	"github.com/findy-network/findy-agent/cmds/agency"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("apns-p12-file", envPrefix+"_AGENCY_APNS_P12_FILE"))
		err2.Check(viper.BindEnv("host-address", envPrefix+"_AGENCY_HOST_ADDRESS"))
		err2.Check(viper.BindEnv("host-port", envPrefix+"_AGENCY_HOST_PORT"))
		err2.Check(viper.BindEnv("server-port", envPrefix+"_AGENCY_SERVER_PORT"))
		err2.Check(viper.BindEnv("service-name", envPrefix+"_AGENCY_SERVICE_NAME"))
		err2.Check(viper.BindEnv("pool-name", envPrefix+"_AGENCY_POOL_NAME"))
		err2.Check(viper.BindEnv("pool-protocol", envPrefix+"_AGENCY_POOL_PROTOCOL"))
		err2.Check(viper.BindEnv("steward-seed", envPrefix+"_AGENCY_STEWARD_SEED"))
		err2.Check(viper.BindEnv("psm-database-file", envPrefix+"_AGENCY_PSM_DATABASE_FILE"))
		err2.Check(viper.BindEnv("reset-register", envPrefix+"_AGENCY_RESET_REGISTER"))
		err2.Check(viper.BindEnv("register-file", envPrefix+"_AGENCY_REGISTER_FILE"))
		err2.Check(viper.BindEnv("steward-wallet-name", envPrefix+"_AGENCY_STEWARD_WALLET_NAME"))
		err2.Check(viper.BindEnv("steward-wallet-key", envPrefix+"_AGENCY_STEWARD_WALLET_KEY"))
		err2.Check(viper.BindEnv("steward-did", envPrefix+"_AGENCY_STEWARD_DID"))
		err2.Check(viper.BindEnv("protocol-path", envPrefix+"_AGENCY_PROTOCOL_PATH"))
		err2.Check(viper.BindEnv("salt", envPrefix+"_AGENCY_SALT"))
		return nil
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
		defer err2.Return(&err)
		err2.Check(viper.BindEnv("base-address", envPrefix+"_AGENCY_PING_BASE_ADDRESS"))
		return nil
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
	flags.StringVar(&aCmd.APNSP12CertFile, "apns-p12-file", "", "APNS certificate p12 file, ENV variable: "+envPrefix+"_AGENCY_APNS_P12_FILE")
	flags.StringVar(&aCmd.HostAddr, "host-address", "localhost", "host address, ENV variable: "+envPrefix+"_AGENCY_HOST_ADDRESS")
	flags.UintVar(&aCmd.HostPort, "host-port", 8080, "host port, ENV variable: "+envPrefix+"_AGENCY_HOST_PORT")
	flags.UintVar(&aCmd.ServerPort, "server-port", 8080, "server port, ENV variable: "+envPrefix+"_AGENCY_SERVER_PORT")
	flags.StringVar(&aCmd.ServiceName, "service-name", "ca-api", "service name, ENV variable: "+envPrefix+"_AGENCY_SERVICE_NAME")
	flags.StringVar(&aCmd.PoolName, "pool-name", "findy-pool", "pool name, ENV variable: "+envPrefix+"_AGENCY_POOL_NAME")
	flags.Uint64Var(&aCmd.PoolProtocol, "pool-protocol", 2, "pool protocol, ENV variable: "+envPrefix+"_AGENCY_POOL_PROTOCOL")
	flags.StringVar(&aCmd.StewardSeed, "steward-seed", "000000000000000000000000Steward1", "steward seed, ENV variable: "+envPrefix+"_AGENCY_STEWARD_SEED")
	flags.StringVar(&aCmd.PsmDb, "psm-database-file", "findy.bolt", "state machine database's filename, ENV variable: "+envPrefix+"_AGENCY_PSM_DATABASE_FILE")
	flags.BoolVar(&aCmd.ResetData, "reset-register", false, "reset handshake register, ENV variable: "+envPrefix+"_AGENCY_RESET_REGISTER")
	flags.StringVar(&aCmd.HandshakeRegister, "register-file", "findy.json", "handshake registry's filename, ENV variable: "+envPrefix+"_AGENCY_REGISTER_FILE")
	flags.StringVar(&aCmd.WalletName, "steward-wallet-name", "", "steward wallet name, ENV variable: "+envPrefix+"_AGENCY_STEWARD_WALLET_NAME")
	flags.StringVar(&aCmd.WalletPwd, "steward-wallet-key", "", "steward wallet key, ENV variable: "+envPrefix+"_AGENCY_STEWARD_WALLET_KEY")
	flags.StringVar(&aCmd.StewardDid, "steward-did", "", "steward DID, ENV variable: "+envPrefix+"_AGENCY_STEWARD_DID")
	flags.StringVar(&aCmd.ServiceName2, "protocol-path", "a2a", "URL path for A2A protocols, ENV variable: "+envPrefix+"_AGENCY_PROTOCOL_PATH") // agency.ProtocolPath is available
	flags.StringVar(&aCmd.Salt, "salt", "", "salt, ENV variable: "+envPrefix+"_AGENCY_SALT")

	p := pingAgencyCmd.Flags()
	p.StringVar(&paCmd.BaseAddr, "base-address", "http://localhost:8080", "base address of agency, ENV variable: "+envPrefix+"_AGENCY_PING_BASE_ADDRESS")

	rootCmd.AddCommand(agencyCmd)
	agencyCmd.AddCommand(startAgencyCmd)
	agencyCmd.AddCommand(pingAgencyCmd)
}
