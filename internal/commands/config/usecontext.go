package config

import (
	"fmt"

	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func newCmdConfigUseContext() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use-context",
		Short:   "Prints the current-context",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: use-context requires a single context name as an argument")
			} else {
				configutils.UseContext(args[0])
			}
		},
	}
	return cmd
}