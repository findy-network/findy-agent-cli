package extracmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/findy-network/findy-agent-cli/cmd"
	"github.com/findy-network/findy-agent-cli/completionhelp"
	"github.com/findy-network/findy-agent/agent/didcomm"
	"github.com/findy-network/findy-agent/cmds/connection"
	"github.com/findy-network/findy-wrapper-go/dto"
	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var issueDoc = `Credentials can issued with the command over existing DIDComm connection`

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Sends the credential offer to other end of the pairwise/connection",
	Long:  issueDoc,
	PreRunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)
		err2.Check(cmd.BindEnvs(envs, "user"))
		return cmd.BindEnvs(proofEnvs, "send")
	},
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		err2.Check(cmdIssue.Validate())
		if !cmd.DryRun() {
			c.SilenceUsage = true

			err2.Try(cmdIssue.Exec(os.Stdout))

			proofAttrs := err2.Try(buildProofAttrs(cmdIssue.Attributes, cmdIssue.CredDefID))[0].([]didcomm.ProofAttribute)
			fmt.Println(dto.ToJSON(proofAttrs))
		}
		return nil

	},
}

func parseAttrs(a string) (credAttrs []didcomm.CredentialAttribute, err error) {
	if err := json.Unmarshal([]byte(a), &credAttrs); err != nil {
		return nil, err
	}
	return credAttrs, nil
}

func buildProofAttrs(a, credDefID string) (proofAttrs []didcomm.ProofAttribute, err error) {
	defer err2.Return(&err)

	credAttrs := err2.Try(parseAttrs(a))[0].([]didcomm.CredentialAttribute)
	proofAttrs = make([]didcomm.ProofAttribute, len(credAttrs))
	for i, attr := range credAttrs {
		proofAttrs[i] = didcomm.ProofAttribute{
			Name:      attr.Name,
			CredDefID: credDefID,
		}
	}
	return proofAttrs, nil
}

var cmdIssue = connection.IssueCmd{}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	issueCmd.Flags().StringVar(&cmdIssue.WalletName, "wallet-name", "", "edge agent's wallet name")
	err2.Check(issueCmd.Flags().SetAnnotation("wallet-name", cobra.BashCompSubdirsInDir, completionhelp.WalletLocations()))

	issueCmd.Flags().StringVar(&cmdIssue.WalletKey, "wallet-key", "", "edge agent's wallet key")
	issueCmd.Flags().StringVarP(&cmdIssue.Name, "conn-id", "i", "", "connection id")
	issueCmd.Flags().StringVar(&cmdIssue.CredDefID, "cred-def-id", "", "cred def id")
	issueCmd.Flags().StringVar(&cmdIssue.Attributes, "attrs", `[{"name":"email","value":"test@email.com"}]`, "credential attributes in JSON")

	cmd.RootCmd().AddCommand(issueCmd)
}

var envs = map[string]string{
	"wallet-name": "WALLET_NAME",
	"wallet-key":  "WALLET_KEY",
}
