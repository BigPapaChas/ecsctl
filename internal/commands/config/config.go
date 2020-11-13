package config

import "github.com/spf13/cobra"

func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Retrieves resources of an ECS cluster",
		DisableFlagsInUseLine: true,
	}

	cmd.AddCommand(newCmdConfigView())
	cmd.AddCommand(newCmdConfigCurrentContext())
	cmd.AddCommand(newCmdConfigUseContext())
	cmd.AddCommand(newCmdConfigGetContext())
	cmd.AddCommand(newCmdConfigCreateContext())
	return cmd
}