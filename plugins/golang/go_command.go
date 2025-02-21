package golang

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"path/filepath"
	"runtime"
	"templarium/plugins/golang/commands/cli"
	"templarium/plugins/sdk"
	"text/template"
)

type GoCommand interface {
	sdk.Command
}

func NewGoCommand(fileSystem afero.Fs) GoCommand {
	return newGoCommand(fileSystem)
}

func newGoCommand(fileSystem afero.Fs) *goCommand {
	goCmd := &goCommand{}

	cobraCommand := &cobra.Command{
		Use:   "go",
		Short: "Create a Go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			moduleName, _ := cmd.Flags().GetString("module-name")

			err := goCmd.run(moduleName, version)
			if err != nil {
				return errors.Wrap(err, "error creating go command")
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
	sdk.BaseCommand
}

func (g goCommand) createGoModFile() (afero.File, error) {
	const goModName = "go.mod"
	file, err := g.FileSystem.Create(goModName)
	if err != nil {
		return nil, errors.Wrap(err, "error creating file")
	}
	return file, nil
}

func (g goCommand) run(moduleName string, version string) error {
	err := validateFlags(moduleName, version)
	if err != nil {
		return err
	}

	file, err := g.createGoModFile()
	defer func(file afero.File) {
		err := file.Close()
		if err != nil {
			panic(errors.Wrap(err, "error closing file"))
		}
	}(file)
	if err != nil {
		return errors.Wrap(err, "error creating go.mod file")
	}

	tmpl, err := template.ParseFiles(filepath.Join(g.getRootDirectory(), "resources/go.mod.tmpl"))
	if err != nil {
		return errors.Wrap(err, "error parsing file")
	}

	err = tmpl.Execute(file, goModVariables{
		GoVersion:  version,
		ModuleName: moduleName,
	})
	if err != nil {
		return errors.Wrap(err, "error executing template")
	}

	return nil
}

func validateFlags(moduleName string, version string) error {
	if version == "" {
		return errors.New("the --version flag is required")
	}
	if moduleName == "" {
		return errors.New("the --module-name flag is required")
	}
	return nil
}

func (g goCommand) getRootDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get root directory")
	}

	return filepath.Dir(filename)
}
