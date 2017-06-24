package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an MFA account",
	Long:  `Remove an MFA account and delete the secret from the keychain`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return remove()
	},
}

func init() {
	accountCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&accountName, "name", "n", "", "A name to identify the MFA account")
}

func remove() error {
	if accountName == "" {
		return errors.New("Missing '--name' in 'remove' command")
	}
	return delete(accountName)
}
