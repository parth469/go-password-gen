/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Generate = &cobra.Command{
	Use:   "generate",
	Short: "Generate a secure password or a batch of passwords with customizable options",
	Long: `The 'generate' command is used to create a secure password or multiple passwords based on customizable parameters like length, inclusion of digits, special characters, or a custom set of characters.

This command gives you flexibility in password creation, ensuring the generated passwords meet specific security requirements or personal preferences.

### Options for password generation:
   - Password length: Use the '-l' flag to specify the length of the password (e.g., -l 12 for a 12-character password).
   - Include numbers: Add digits by using the '-d' flag.
   - Include special characters: Use the '-s' flag to include special characters.
   - Custom character set: Provide a custom string of characters using the '-c' flag.
   - Bulk generation: Generate multiple passwords at once with the '-b' flag, followed by the number of passwords to generate.

### Example usage:
1) Generate a 12-character password with numbers and special characters:
   password generate -l 12 -d -s 

2) Generate 10 passwords with a length of 16 characters:
   password generate -b 10 -l 16

This command is designed to offer robust security through customizable password generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}
