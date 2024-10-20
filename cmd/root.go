/*
Copyright © 2024 parth
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "password",
	Short: "A command-line tool for generating and managing secure passwords",
	Long: `Password is a customizable CLI application built in Go, designed to generate secure passwords and manage them.

You can specify the following options for password generation and management:
	
1) Password generation:
   - Password length: Choose the desired length of the password using the '-l' flag (e.g., -l 12 for a 12-character password).
   - Include numbers: Add digits to the password by using the flag '-d'.
   - Include special characters: Add special characters by using the flag '-s'.
   - Custom string: Provide a specific set of characters to generate the password from using the '-c' flag.
   - Bulk generation: Generate multiple passwords at once using the '-b' flag followed by the number of passwords.

2) Password management:
   - Save passwords: Use the 'save' option to store passwords for future use, along with additional details such as purpose and creation date.
   - List passwords: Use  to list all saved passwords, or use the '-a' flag to display all generated passwords, even if unsaved.
   - Update saved passwords: Use to update the purpose of a saved password.
   - Delete passwords: Use  to delete a saved password based on its ID.

### Example usage:
1) Generate a password with length 12, including numbers and special characters:
	password generate -l 12 -d -s 

2) Generate a bulk of 10 passwords with a custom length of 16:
	password generate -b 10 -l 16

3) Save a password with a purpose:
	password save -g -l 12 -d -s --purpose "WiFi password"

4) List all saved passwords:
	password list

5) Delete a password by ID:
	password delete <id> 

5) Update a password by ID:
	password update <id> 


This tool provides flexibility for generating secure passwords with detailed management features for saving, listing, updating, and deleting them.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
