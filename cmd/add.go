package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var force bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new MFA account",
	Long:  `Add a new MFA account by providing the secret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return add()
	},
}

func init() {
	accountCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&accountName, "name", "n", "", "A name to identify the MFA account")
	addCmd.Flags().StringVarP(&accountSecret, "secret", "s", "", "A secret used to register the MFA account")
	addCmd.Flags().BoolVarP(&force, "force", "f", false, "Force an account secret to be overwritten")
}

func add() error {
	if accountName == "" {
		return errors.New("Missing '--name' in 'add' command")
	}
	if accountSecret == "" {
		return errors.New("Missing '--secret' in 'add' command")
	}
	return save(accountName, accountSecret)
}
