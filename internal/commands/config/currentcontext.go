package config

import (
	"fmt"
	"os"

	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func newCmdConfigCurrentContext() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "current-context",
		Short:   "Prints the current-context",
		Run: func(cmd *cobra.Command, args []string) {
			if err := configutils.PrintCurrentContext(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return cmd
}