package cli

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"templarium/plugins/golang/codegen"
	"templarium/runner"
	sdk2 "templarium/sdk"
)

type CliCommand interface {
	sdk2.Command
}

func NewCliCommand(fileSystem afero.Fs) CliCommand {
	return newCliCommand(fileSystem)
}

func newCliCommand(fileSystem afero.Fs) *cliCommand {
	cliCmd := &cliCommand{}

	cobraCommand := &cobra.Command{
		Use:   "cli",
		Short: "Create a CLI Go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			moduleName, _ := cmd.Flags().GetString("module-name")

			err := codegen.NewGoCodeGenerator(fileSystem, runner.DefaultRunner()).CreateGoCliProject(moduleName, version)
			if err != nil {
				return errors.Wrap(err, "error creating go project")
			}

			return nil
		},
	}

	cobraCommand.Flags().StringP("version", "v", "", "Go project version")
	cobraCommand.Flags().StringP("module-name", "m", "", "Module name")

	cliCmd.CobraCommand = cobraCommand
	cliCmd.FileSystem = fileSystem

	return cliCmd
}

type cliCommand struct {
	sdk2.BaseCommand
}
