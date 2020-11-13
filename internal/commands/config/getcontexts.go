package config

import (
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func newCmdConfigGetContext() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get-contexts",
		Short:   "Prints the available contexts",
		Run: func(cmd *cobra.Command, args []string) {
			configutils.PrintContexts()
		},
	}
	return cmd
}