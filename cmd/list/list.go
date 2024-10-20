/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "List all saved passwords or generated passwords",
	Long: `The 'list' command displays all passwords saved in the management system. You can also list passwords generated during the current session, even if they were not saved, by using the '-a' flag.

This command is helpful for reviewing the passwords you've generated or saved, making it easy to track which passwords are stored and for what purpose.

### Example usage:
1) List all saved passwords:
   password list

2) List all passwords generated in the session, even if unsaved:
   password list -a

This command provides quick access to your password inventory, helping you manage your saved and generated passwords effectively.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}
