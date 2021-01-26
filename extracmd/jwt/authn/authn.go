package authn

import (
	"fmt"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/extracmd/jwt"
	"github.com/findy-network/findy-grpc/acator/authn"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var acatorDoc = `Authn is headless WebAuthn authenticator.

The authenticator allows Register and Login to WebAuthn server.`

var acatorCmd = &cobra.Command{
	Use:   "authn",
	Short: "authn commands",
	Long:  acatorDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		return cmd.BindEnvs(envs, "authn")
	},
	Run: func(c *cobra.Command, args []string) {
		cmd.SubCmdNeeded(c)
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
