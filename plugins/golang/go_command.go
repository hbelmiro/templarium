package golang

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

var command *cobra.Command

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

			createGoModFile(moduleName, version)

			fmt.Printf("Go version %s!\n", version)
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

func createGoModFile(moduleName string, version string) {
	file, err := os.Create("go.mod-test")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get current file path")
	}

	baseDir := filepath.Dir(filename)

	tmpl, err := template.ParseFiles(filepath.Join(baseDir, "go.mod"))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, Data{
		GoVersion:  version,
		ModuleName: moduleName,
	})
	if err != nil {
		panic(err)
	}
}

func Command() *cobra.Command {
	return command
}
