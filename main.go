package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

func main() {
	var rootCmd = &cobra.Command{Use: "TúCLI"}
	var name, email, password string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "You can create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				fmt.Println("Name can't be empty")
				return
			}
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
			if !emailRegex.MatchString(email) {
				fmt.Println("Invalid email, try a valid one")
				return
			}

			if len(password) < 6 {
				fmt.Println("The password must be at least 6 characters")
				return
			}
			fmt.Printf("Name %s\nEmail: %s\nPasswor: %s\n", name, email, password)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "User name")
	cmd.Flags().StringVarP(&email, "email", "e", "", "User email")
	cmd.Flags().StringVarP(&password, "password", "p", "", "User password")

	rootCmd.AddCommand(cmd)
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
