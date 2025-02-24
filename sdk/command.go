package sdk

import (
	"github.com/spf13/cobra"
)

type Command interface {
	GetCobraCommand() *cobra.Command
}
