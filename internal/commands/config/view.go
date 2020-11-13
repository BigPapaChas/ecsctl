package config

import (
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func newCmdConfigView() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "view",
		Short:   "View config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return configutils.PrintConfig()
		},
	}
	return cmd
}