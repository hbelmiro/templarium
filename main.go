package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"templarium/plugins/golang"
)

func main() {
	// templarium add go --go-version=1.23.5
	var rootCmd = &cobra.Command{
		Use:   "templarium",
		Short: "A simple CLI tool",
	}

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greet a user",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	greetCmd.Flags().String("name", "World", "Name to greet")
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(golang.Command())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
