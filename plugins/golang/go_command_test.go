package golang

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGoCommand(t *testing.T) {
	fileSystem := afero.NewMemMapFs()

	err := newGoCommand(fileSystem).run("my-module", "1.22")
	require.NoError(t, err)

	file, err := afero.ReadFile(fileSystem, "go.mod")
	require.NotNil(t, file)

	expectedContents, err := os.ReadFile("test-resources/go.mod-test")
	require.NoError(t, err)
	require.Equal(t, string(expectedContents), string(file))
}
