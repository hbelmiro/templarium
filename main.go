package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
