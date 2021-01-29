package authn

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/acator/authn"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var acatorDoc = `Authn is headless WebAuthn authenticator.

The authenticator allows Register and Login to WebAuthn server. When prefilled
JSON cmd is sent thru stdio or file, it's treated as secondary data source. That
means that any its dictionary item can be over written by command flag values'`

var acatorCmd = &cobra.Command{
	Use:   "authn",
	Short: "authn commands",
	Long:  acatorDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "authn")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		if len(args) == 0 {
			return errors.New("input missing")
		}

		inJSON := os.Stdin
		if args[0] != "-" {
			inJSON = err2.File.Try(os.Open(args[0]))
			defer inJSON.Close()
		}
		execCmd := authnCmd.TryReadJSON(inJSON)

		if !cmd.DryRun() {
			var r authn.Result
			r, err = execCmd.Exec(os.Stdout)
			err2.Check(err)
			fmt.Println(r.String())
		} else {
			b, _ := json.MarshalIndent(execCmd, "", "  ")
			fmt.Println(string(b))
		}

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	acatorCmd.PersistentFlags().StringVarP(&authnCmd.UserName, "user_name", "u", "", "used for registration name")
	acatorCmd.PersistentFlags().StringVar(&authnCmd.Url, "url", authnCmd.Url, "WebAuthn server URL aka origin")
	acatorCmd.PersistentFlags().StringVar(&authnCmd.Key, "key", authnCmd.Key, "master key for authenticator")
	acatorCmd.PersistentFlags().StringVar(&authnCmd.AAGUID, "aaguid", authnCmd.AAGUID, "authenticator AAGUID")
	acatorCmd.PersistentFlags().Uint64Var(&authnCmd.Counter, "counter", authnCmd.Counter, "authenticator counter")

	jwt.JwtCmd.AddCommand(acatorCmd)
}

var authnCmd = authn.Cmd{
	SubCmd:   "",
	UserName: "",
	Url:      "http://localhost:8090",
	AAGUID:   "12c85a48-4baf-47bd-b51f-f192871a1511",
	Key:      "15308490f1e4026284594dd08d31291bc8ef2aeac730d0daf6ff87bb92d4336c",
	Counter:  0,
}

var envs = map[string]string{
	"url":     "URL",
	"aaguid":  "AAGUID",
	"key":     "KEY",
	"counter": "COUNTER",
}
