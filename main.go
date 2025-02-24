package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
	"templarium/plugins/golang"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "templarium",
		Short: "This is a CLI application for generating project templates.",
	}

	fileSystem := afero.NewOsFs()

	rootCmd.AddCommand(golang.NewGoCommand(fileSystem).GetCobraCommand())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
