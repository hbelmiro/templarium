package main

import (
	"github.com/spf13/afero"
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

	fileSystem := afero.NewOsFs()

	goCommand := golang.NewGoCommand(fileSystem)

	rootCmd.AddCommand(goCommand.GetCobraCommand())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
