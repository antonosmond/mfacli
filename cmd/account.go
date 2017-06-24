package cmd

import "github.com/spf13/cobra"

var accountName string
var accountSecret string

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Add, remove, or list accounts",
	Long:  `Add, remove, or list accounts`,
}

func init() {
	RootCmd.AddCommand(accountCmd)
}
