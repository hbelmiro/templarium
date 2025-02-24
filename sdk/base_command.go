package sdk

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type BaseCommand struct {
	FileSystem   afero.Fs
	CobraCommand *cobra.Command
}

func (c BaseCommand) GetCobraCommand() *cobra.Command {
	return c.CobraCommand
}

func (c BaseCommand) AddCommand(command Command) {
	c.CobraCommand.AddCommand(command.GetCobraCommand())
}
