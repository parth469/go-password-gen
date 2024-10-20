/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Update = &cobra.Command{
	Use:   "update",
	Short: "Update the purpose of a saved password by ID",
	Long: `The 'update' command allows you to modify the details of a saved password, such as its purpose. 
For example, if a saved password's original purpose has changed, you can use this command to update 
that information in your password management system.

You can specify the ID of the saved password you wish to update and provide the new purpose, ensuring your password management system stays accurate and up to date.

### Example usage:

1) Update the purpose of a password:
   password update <id> --purpose "New purpose"

This command is particularly useful when you need to track changes in the usage of passwords without generating new ones.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
