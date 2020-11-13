package config

import (
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func newCmdConfigCreateContext() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-context",
		Short:   "Create a new context",
		RunE: func(cmd *cobra.Command, args []string) error {
			return configutils.CreateNewContext()
		},
	}
	return cmd
}