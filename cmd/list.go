package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the accounts",
	Long:  `List the MFA accounts`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return list()
	},
}

func init() {
	accountCmd.AddCommand(listCmd)
}

func list() error {
	accounts, err := query()
	if err != nil {
		return err
	}
	for _, a := range accounts {
		fmt.Println(a)
	}
	return nil
}
