package cmd

import (
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(bitbucket completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(bitbucket completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletionFile("scripts/bash_completion.sh")
		rootCmd.GenPowerShellCompletionFile("scripts/powershell_completion.sh")
		rootCmd.GenZshCompletionFile("scripts/zsh_completion.sh")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
