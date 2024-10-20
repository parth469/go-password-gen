/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a saved password by its ID",
	Long: `The 'delete' command removes a saved password from the password management system based on its unique ID.

This command is essential for managing your stored passwords, allowing you to clean up passwords that are no longer in use or needed for security reasons.

### Example usage:
1) Delete a password by providing its ID:
   password delete <id>

This helps in keeping your password management system organized by removing outdated or unnecessary entries.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}
