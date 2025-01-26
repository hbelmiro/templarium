package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
	"templarium/plugins/golang"
)

func main() {
	// go run main.go go --version=123 --module-name=Helber
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
