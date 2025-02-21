package cli

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"templarium/plugins/sdk"
)

type CliCommand interface {
	sdk.Command
}

func NewCliCommand(fileSystem afero.Fs) CliCommand {
	return newCliCommand(fileSystem)
}

func newCliCommand(fileSystem afero.Fs) *cliCommand {
	c := &cliCommand{}

	howdyCommand := &cobra.Command{
		Use:   "howdy",
		Short: "say howdy",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			print("Howdy " + name)
			return nil
		},
	}

	howdyCommand.Flags().StringP("name", "n", "", "Howdy name")

	c.CobraCommand = howdyCommand
	c.FileSystem = fileSystem

	return c
}

type cliCommand struct {
	sdk.BaseCommand
}
