package cmd

import (
	goflag "flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/findy-network/findy-agent-cli/utils"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/findy-network/findy-common-go/rpc"
	"github.com/golang/glog"
	"github.com/lainio/err2"
	"github.com/lainio/err2/assert"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const envPrefix = "FCLI"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "findy-agent-cli",
	Short: "Findy agent cli tool",
	Long:  `Findy agent cli tool`,

	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
		defer err2.Handle(&err)

		// NOTE! Very important. Adds support for std flag pkg users: glog, err2
		goflag.Parse()

		// let's support our old logging env
		use, lvl := utils.ParseLoggingArgs(rootFlags.logging)
		if use {
			try.To(goflag.Set("v", strconv.Itoa(lvl)))
		}
		try.To(goflag.Set("logtostderr", "true"))
		handleViperFlags(cmd)
		glog.CopyStandardLogTo("ERROR") // for err2
		glog.V(13).Infoln("flag inits OK")
		return nil
	},
}

// Execute root
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// To fix errors printed twice removing the cobra generators next
		// see: https://github.com/spf13/cobra/issues/304
		// fmt.Println(err)

		os.Exit(1)
	}
}

// RootCmd returns a current root command which can be used for adding own
// commands in an own repo.
//
//	implCmd.AddCommand(listCmd)
//
// That's a helper function to extend this CLI with own commands and offering
// same base commands as this CLI.
func RootCmd() *cobra.Command {
	return rootCmd
}

// DryRun returns a value of a dry run flag. That's a helper function to extend
// this CLI with own commands and offering same base commands as this CLI.
func DryRun() bool {
	return rootFlags.dryRun
}

func ServiceAddr() string {
	return rootFlags.ServiceAddr
}

func TLSPath() string {
	return rootFlags.TLSPath
}

func BaseCfg() (_ *rpc.ClientCfg, err error) {
	defer err2.Handle(&err)

	var baseCfg *rpc.ClientCfg
	if TLSPath() != "" {
		baseCfg = client.BuildConnBase(TLSPath(), ServiceAddr(), nil)
	} else {
		var (
			addr string
			port int
		)
		s := ServiceAddr()
		glog.V(15).Infoln("ServiceAddr string from flag:", s)
		ss := strings.Split(s, ":")
		assert.SLen(ss, 2)
		addr = ss[0]
		port = try.To1(strconv.Atoi(ss[1]))
		glog.V(3).Infoln("ServiceAddr:", addr, port)

		baseCfg = client.BuildInsecureClientConnBase(addr, port, nil)
	}
	return baseCfg, nil
}

// RootFlags are the common flags
type RootFlags struct {
	cfgFile     string
	dryRun      bool
	logging     string
	ServiceAddr string
	TLSPath     string
}

// ClientFlags agent flags
type ClientFlags struct {
	WalletName string
	WalletKey  string
	URL        string
}

var rootFlags = RootFlags{}

var rootEnvs = map[string]string{
	"config":   "CONFIG",
	"logging":  "CLI_LOGGING",
	"dry-run":  "DRY_RUN",
	"server":   "SERVER",
	"tls-path": "TLS_PATH",
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		log.Println(err)
	}))

	cobra.OnInitialize(initConfig)

	flags := rootCmd.PersistentFlags()
	flags.StringVar(&rootFlags.cfgFile, "config", "",
		FlagInfo("configuration file", "", rootEnvs["config"]))
	flags.StringVar(&rootFlags.logging, "logging", "",
		FlagInfo("logging startup arguments", "", rootEnvs["logging"]))
	flags.StringVar(&rootFlags.ServiceAddr, "server", "localhost:50051",
		FlagInfo("gRPC server addr:port", "", rootEnvs["server"]))
	flags.StringVar(&rootFlags.TLSPath, "tls-path", "",
		FlagInfo("TLS cert path", "", rootEnvs["tls-path"]))
	flags.BoolVarP(&rootFlags.dryRun, "dry-run", "n", false,
		FlagInfo("perform a trial run with no changes made", "", rootEnvs["dry-run"]))

	try.To(viper.BindPFlag("logging", flags.Lookup("logging")))
	try.To(viper.BindPFlag("dry-run", flags.Lookup("dry-run")))
	try.To(viper.BindPFlag("server", flags.Lookup("server")))
	try.To(viper.BindPFlag("tls-path", flags.Lookup("tls-path")))

	BindEnvs(rootEnvs, "")

	// NOTE! Very important. Adds support for std flag pkg users: glog, err2
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func initConfig() {
	viper.SetEnvPrefix(envPrefix)
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	readConfigFile()
	readBoundRootFlags()
}

func readBoundRootFlags() {
	rootFlags.logging = viper.GetString("logging")
	rootFlags.dryRun = viper.GetBool("dry-run")
}

func readConfigFile() {
	cfgEnv := os.Getenv(getEnvName("", "config"))
	if rootFlags.cfgFile != "" || cfgEnv != "" {
		printInfo := true
		if rootFlags.cfgFile == "" {
			rootFlags.cfgFile = cfgEnv
			printInfo = false
		}
		viper.SetConfigFile(rootFlags.cfgFile)
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil && printInfo {
			glog.V(3).Infoln("Using config file:", viper.ConfigFileUsed())
		}
	}
}

// BindEnvs calls viper.BindEnv with envMap and cmdName which can be empty if
// flag is general.
func BindEnvs(envMap map[string]string, cmdName string) (err error) {
	defer err2.Handle(&err)
	for flagKey, envName := range envMap {
		finalEnvName := getEnvName(cmdName, envName)
		try.To(viper.BindEnv(flagKey, finalEnvName))
	}
	return nil
}

func FlagInfo(info, cmdPrefix, envName string) string {
	return info + ", " + getEnvName(cmdPrefix, envName)
}

func getEnvName(cmdName, envName string) string {
	if cmdName == "" {
		return envPrefix + "_" + strings.ToUpper(envName)
	}
	return envPrefix + "_" + strings.ToUpper(cmdName) + "_" + envName
}

func handleViperFlags(cmd *cobra.Command) {
	setRequiredStringFlags(cmd)
	if cmd.HasParent() {
		handleViperFlags(cmd.Parent())
	}
}

func setRequiredStringFlags(cmd *cobra.Command) {
	viper.BindPFlags(cmd.LocalFlags())
	if cmd.PreRunE != nil {
		cmd.PreRunE(cmd, nil)
	}
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		if viper.GetString(f.Name) != "" {
			cmd.LocalFlags().Set(f.Name, viper.GetString(f.Name))
		}
	})
}

// SubCmdNeeded prints the help and error messages because the cmd is abstract.
func SubCmdNeeded(cmd *cobra.Command) {
	fmt.Println("Subcommand needed!")
	cmd.Help()
	os.Exit(1)
}
