package golang

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"templarium/plugins/golang/codegen"
	"templarium/plugins/golang/commands/cli"
	"templarium/runner"
	sdk2 "templarium/sdk"
)

type GoCommand interface {
	sdk2.Command
}

func NewGoCommand(fileSystem afero.Fs) GoCommand {
	return newGoCommand(fileSystem)
}

func newGoCommand(fileSystem afero.Fs) *goCommand {
	goCmd := &goCommand{
		goCodeGenerator: codegen.NewGoCodeGenerator(fileSystem, runner.DefaultRunner()),
	}

	cobraCommand := &cobra.Command{
		Use:   "go",
		Short: "Create a Go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			moduleName, _ := cmd.Flags().GetString("module-name")

			err := goCmd.goCodeGenerator.CreateGoProject(moduleName, version)
			if err != nil {
				return errors.Wrap(err, "error creating go project")
			}

			return nil
		},
	}

	cobraCommand.Flags().StringP("version", "v", "", "Go project version")
	cobraCommand.Flags().StringP("module-name", "m", "", "Module name")

	goCmd.CobraCommand = cobraCommand
	goCmd.FileSystem = fileSystem

	goCmd.AddCommand(cli.NewCliCommand(goCmd.FileSystem))

	return goCmd
}

type goCommand struct {
	sdk2.BaseCommand

	goCodeGenerator codegen.GoCodeGenerator
}
