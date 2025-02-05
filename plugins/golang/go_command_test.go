package golang

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGoCommand(t *testing.T) {
	fileSystem := afero.NewMemMapFs()

	err := newGoCommand(fileSystem).run("my-module", "1.22")
	require.NoError(t, err)

	file, err := afero.ReadFile(fileSystem, "go.mod-test")
	require.NotNil(t, file)

	expectedContents := "module my-module\n\ngo 1.22"
	require.Equal(t, expectedContents, string(file))
}
