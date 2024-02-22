package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/spf13/cobra"
)

var newKeyDoc = ``

var newKeyCmd = &cobra.Command{
	Use:   "new-key",
	Short: "Create a new key for the authenticator",
	Long:  newKeyDoc,
	RunE: func(*cobra.Command, []string) (err error) {
		defer err2.Handle(&err)

		key := make([]byte, 32)
		try.To1(rand.Read(key))
		fmt.Println(hex.EncodeToString(key))

		return nil
	},
}

func init() {
	defer err2.Catch(err2.Err(func(err error) {
		fmt.Println(err)
	}))

	rootCmd.AddCommand(newKeyCmd)
}
