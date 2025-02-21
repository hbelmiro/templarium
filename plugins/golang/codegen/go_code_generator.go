package codegen

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"path/filepath"
	"runtime"
	"strings"
	"templarium/runner"
	"text/template"
)

type GoCodeGenerator interface {
	CreateGoProject(moduleName string, version string) error
	CreateGoCliProject(moduleName string, version string) error
}

func NewGoCodeGenerator(fileSystem afero.Fs, cobraVersionRetriever runner.Runner) GoCodeGenerator {
	return &goCodeGenerator{
		fileSystem:            fileSystem,
		cobraVersionRetriever: cobraVersionRetriever,
	}
}

type goCodeGenerator struct {
	fileSystem            afero.Fs
	cobraVersionRetriever runner.Runner
}

func (g goCodeGenerator) CreateGoProject(moduleName string, version string) error {
	return g.createGoProject(moduleName, version, "resources/go.mod.tmpl", &goModVariables{
		GoVersion:  version,
		ModuleName: moduleName,
	})
}

func (g goCodeGenerator) CreateGoCliProject(moduleName string, version string) error {
	cobraVersion, err := g.getLatestCobraVersion()
	if err != nil {
		return errors.Wrap(err, "error creating Go CLI project")
	}

	variables := goModVariables{
		GoVersion:    version,
		ModuleName:   moduleName,
		CobraVersion: cobraVersion,
	}

	return g.createGoProject(moduleName, version, "resources/go.mod-cli.tmpl", &variables)
}

func (g goCodeGenerator) createGoProject(moduleName string, version string, goModtemplatePath string, variables *goModVariables) error {
	err := validateFlags(moduleName, version)
	if err != nil {
		return errors.Wrap(err, "error creating Go project")
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

	tmpl, err := template.ParseFiles(filepath.Join(g.getRootDirectory(), goModtemplatePath))
	if err != nil {
		return errors.Wrap(err, "error parsing file")
	}

	err = tmpl.Execute(file, variables)
	if err != nil {
		return errors.Wrap(err, "error executing template")
	}

	return nil
}

func (g goCodeGenerator) getLatestCobraVersion() (string, error) {
	output, err := g.cobraVersionRetriever.RunCommand("go", "list", "-m", "-versions", "github.com/spf13/cobra")
	if err != nil {
		return "", errors.Wrap(err, "error getting latest cobra version")
	}

	versions := strings.Fields(string(output))
	latestVersion := versions[len(versions)-1]
	return latestVersion, nil
}

func (g goCodeGenerator) getRootDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get root directory")
	}

	return filepath.Dir(filename)
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

func (g goCodeGenerator) createGoModFile() (afero.File, error) {
	const goModName = "go.mod"
	file, err := g.fileSystem.Create(goModName)
	if err != nil {
		return nil, errors.Wrap(err, "error creating file")
	}
	return file, nil
}
