package golang

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var command *cobra.Command

func init() {
	command = &cobra.Command{
		Use:   "go",
		Short: "Create a Go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, _ := cmd.Flags().GetString("version")
			if version == "" {
				return errors.New("the --version flag is required")
			}
			fmt.Printf("Go version %s!\n", version)
			return nil
		},
	}

	command.Flags().StringP("version", "v", "", "Go project version")
}

func Command() *cobra.Command {
	return command
}
