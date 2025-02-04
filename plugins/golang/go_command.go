package golang

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

var command *cobra.Command

const goModName = "go.mod-test"

func init() {
	command = &cobra.Command{
		Use:   "go",
		Short: "Create a Go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			moduleName, _ := cmd.Flags().GetString("module-name")

			if version == "" {
				return errors.New("the --version flag is required")
			}

			if moduleName == "" {
				return errors.New("the --module-name flag is required")
			}

			err := createGoModFile(moduleName, version)
			if err != nil {
				return errors.Wrap(err, "error creating go.mod file")
			}

			return nil
		},
	}

	command.Flags().StringP("version", "v", "", "Go project version")
	command.Flags().StringP("module-name", "m", "", "Module name")
}

type Data struct {
	GoVersion  string
	ModuleName string
}

func createGoModFile(moduleName string, version string) error {
	file, err := os.Create(goModName)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(errors.Wrap(err, "error closing file"))
		}
	}(file)

	tmpl, err := template.ParseFiles(filepath.Join(getRootDirectory(), "go.mod.tmpl"))
	if err != nil {
		return errors.Wrap(err, "error parsing file")
	}

	err = tmpl.Execute(file, Data{
		GoVersion:  version,
		ModuleName: moduleName,
	})
	if err != nil {
		return errors.Wrap(err, "error executing template")
	}

	return nil
}

func getRootDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get root directory")
	}

	return filepath.Dir(filename)
}

func Command() *cobra.Command {
	return command
}
