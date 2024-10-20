/*
Copyright © 2024 parth
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "password-gen",
	Short: "A command-line tool for generating and managing secure passwords",
	Long: `Password-gen is a versatile and customizable CLI application built in Go, designed to generate secure passwords and manage them.

You can specify the following options for password generation and management:
	
1) Password generation:
   - Password length: Choose the desired length of the password using the '-l' flag (e.g., -l 12 for a 12-character password).
   - Include numbers: Add digits to the password by using the flag '-d'.
   - Include special characters: Add special characters by using the flag '-s'.
   - Custom string: Provide a specific set of characters to generate the password from using the '-c' flag.
   - Bulk generation: Generate multiple passwords at once using the '-b' flag followed by the number of passwords.

2) Password management:
   - Save passwords: Use the 'save' option to store passwords for future use, along with additional details such as purpose and creation date.
   - List passwords: Use the '-h' flag to list all saved passwords, or use the '-a' flag to display all generated passwords, even if unsaved.
   - Update saved passwords: Use the '-u' flag to update the purpose of a saved password.
   - Delete passwords: Use the '-d' flag to delete a saved password based on its ID.

### Example usage:
1) Generate a password with length 12, including numbers and special characters:
	password-gen -g -l 12 -d -s 

2) Generate a bulk of 10 passwords with a custom length of 16:
	password-gen -g -b 10 -l 16

3) Save a password with a purpose:
	password-gen save -g -l 12 -d -s --purpose "WiFi password"

4) List all saved passwords:
	password-gen -h

5) Delete a password by ID:
	password-gen -d <id> 

This tool provides flexibility for generating secure passwords with detailed management features for saving, listing, updating, and deleting them.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.password.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
