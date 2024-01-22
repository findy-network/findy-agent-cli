package authn

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-auth/acator/authn"
	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var acatorDoc = `Authn is headless WebAuthn authenticator.

The authenticator allows Register and Login to WebAuthn server. When prefilled
JSON cmd is sent thru stdio or file, it's treated as secondary data source. That
means that any its dictionary item can be over written by command flag values'`

var acatorCmd = &cobra.Command{
	Use:   "authn",
	Short: "WebAuthn commands",
	Long:  acatorDoc,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Handle(&err)

		inJSON := os.Stdin
		if args[0] != "-" {
			inJSON = try.To1(os.Open(args[0]))
			defer inJSON.Close()
		}
		execCmd := authnCmd.TryReadJSON(inJSON)

		if cmd.DryRun() {
			b, _ := json.MarshalIndent(execCmd, "", "  ")
			fmt.Println(string(b))
			return nil
		}
		r := try.To1(execCmd.Exec(os.Stdout))
		fmt.Println(r.String())

		return nil
	},
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	flags := acatorCmd.PersistentFlags()
	flags.StringVarP(&authnCmd.UserName, "user-name", "u", "",
		cmd.FlagInfo("used for registration name", "", envs["user-name"]))
	flags.StringVarP(&authnCmd.PublicDIDSeed, "seed", "", "",
		cmd.FlagInfo("public DID seed for registration", "", envs["seed"]))
	flags.StringVar(&authnCmd.URL, "url", authnCmd.URL,
		cmd.FlagInfo("WebAuthn server connection URL", "", envs["url"]))
	flags.StringVar(&authnCmd.Origin, "origin", authnCmd.Origin,
		cmd.FlagInfo("Different Origin to use, see --url", "", envs["origin"]))
	flags.StringVar(&authnCmd.Key, "key", authnCmd.Key,
		cmd.FlagInfo("master key for authenticator", "", envs["key"]))
	flags.StringVar(&authnCmd.AAGUID, "aaguid", authnCmd.AAGUID,
		cmd.FlagInfo("authenticator AAGUID", "", envs["aaguid"]))
	flags.Uint64Var(&authnCmd.Counter, "counter", authnCmd.Counter,
		cmd.FlagInfo("authenticator counter", "", envs["counter"]))
	flags.BoolVar(&authnCmd.Legacy, "legacy", authnCmd.Legacy,
		cmd.FlagInfo("authenticator legacy", "", envs["legacy"]))

	acatorCmd.MarkPersistentFlagRequired("url")
	acatorCmd.MarkPersistentFlagRequired("aaguid")
	acatorCmd.MarkPersistentFlagRequired("counter")
	acatorCmd.MarkPersistentFlagRequired("key")

	cmd.RootCmd().AddCommand(acatorCmd)
}

var (
	authnCmd = authn.Cmd{
		SubCmd:        "",
		UserName:      "",
		PublicDIDSeed: "",
		URL:           "http://localhost:8090",
		AAGUID:        "12c85a48-4baf-47bd-b51f-f192871a1511",
		Key:           "",
		Counter:       0,
		Token:         "",
	}

	envs = map[string]string{
		"url":       "URL",
		"aaguid":    "AAGUID",
		"key":       "KEY",
		"counter":   "COUNTER",
		"jwt":       "JWT",
		"origin":    "ORIGIN",
		"user-name": "USER",
		"seed":      "SEED",
	}
)
