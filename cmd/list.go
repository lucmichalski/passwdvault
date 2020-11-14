package cmd

import (
	"fmt"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all password identifiers available",
	Long:  "examples here...",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := configuration.CheckInitFile()
		if err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := db.GetAllKeys()
		if err != nil {
			fmt.Println("Could not print all keys")
		}

		for _, key := range keys {
			fmt.Println(string(key))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
