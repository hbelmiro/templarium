package codegen

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"os"
	"templarium/runner"
	"testing"
)

func TestCreateGoProject(t *testing.T) {
	fileSystem := afero.NewMemMapFs()

	err := NewGoCodeGenerator(fileSystem, runner.DefaultRunner()).CreateGoProject("my-module", "1.22")
	require.NoError(t, err)

	file, err := afero.ReadFile(fileSystem, "go.mod")
	require.NotNil(t, file)

	expectedContents, err := os.ReadFile("test-resources/go.mod-test-create-go-project")
	require.NoError(t, err)
	require.Equal(t, string(expectedContents), string(file))
}

func TestCreateGoCliProject(t *testing.T) {
	fileSystem := afero.NewMemMapFs()

	err := NewGoCodeGenerator(fileSystem, runner.FakeRunnerReturning("github.com/spf13/cobra v1.8.1 v1.9.0 v1.9.1")).
		CreateGoCliProject("my-module", "1.22")
	require.NoError(t, err)

	goModFile, err := afero.ReadFile(fileSystem, "go.mod")
	require.NotNil(t, goModFile)

	goModFileExpectedContents, err := os.ReadFile("test-resources/go.mod-test-create-go-cli-project")
	require.NoError(t, err)
	require.Equal(t, string(goModFileExpectedContents), string(goModFile))

	mainFile, err := afero.ReadFile(fileSystem, "main.go")
	require.NotNil(t, mainFile)

	mainFileExpectedContents, err := os.ReadFile("resources/cli-main.tmpl")
	require.NoError(t, err)
	require.Equal(t, string(mainFileExpectedContents), string(mainFile))
}
