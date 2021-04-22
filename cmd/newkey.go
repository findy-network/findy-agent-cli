package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/lainio/err2"
	"github.com/spf13/cobra"
)

var newKeyDoc = ``

var newKeyCmd = &cobra.Command{
	Use:   "new-key",
	Short: "Create a new key for the authenticator",
	Long:  newKeyDoc,
	RunE: func(c *cobra.Command, args []string) (err error) {
		defer err2.Return(&err)

		key := make([]byte, 32)
		err2.Try(rand.Read(key))
		fmt.Println(hex.EncodeToString(key))

		return nil
	},
}

func init() {
	defer err2.Catch(func(err error) {
		fmt.Println(err)
	})

	rootCmd.AddCommand(newKeyCmd)
}
