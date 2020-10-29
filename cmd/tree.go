package cmd

import (
	"fmt"

	. "github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var treeDoc = `Prints the findy-agent-cli command structure.

The whole command structure is printed if no argument is given.
If command name is given as argument, only specified command structure is printed.
(Command must be direct subcommand of the root command.)
`

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Prints the findy-agent-cli command structure",
	Long:  treeDoc,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer Return(&err)
		if len(args) == 0 {
			printStructure(rootCmd, 0)
		} else {
			c, _, e := rootCmd.Find(args)
			Check(e)
			printStructure(c, 0)
		}
		return nil
	},
}

func printStructure(cmd *cobra.Command, spaces int) {
	for i := spaces; i > 0; i-- {
		fmt.Print("    ")
		if i != 1 {
			fmt.Print("│")
		}
	}
	fmt.Print("├──")
	fmt.Println(cmd.Name())
	for _, subCmd := range cmd.Commands() {
		printStructure(subCmd, spaces+1)
	}
}
func init() {
	rootCmd.AddCommand(treeCmd)
}
