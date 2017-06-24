package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a code for a given MFA account",
	Long:  `Get a code for a given MFA account`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return get()
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&accountName, "name", "n", "", "A name to identify the MFA account")
}

func get() error {
	if accountName == "" {
		return errors.New("Missing '--name' in 'get' command")
	}
	secret, err := load(accountName)
	if err != nil {
		return err
	}
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return err
	}
	fmt.Printf(code)
	return nil
}
