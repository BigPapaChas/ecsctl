package configutils

import (
	"github.com/spf13/cobra"
)

func AddCommonPersistentFlags(cmd *cobra.Command) {
	// Create persistent flags on passed command
	cmd.PersistentFlags().StringP("region", "r", "us-east-1", "AWS region")
	cmd.PersistentFlags().StringP("profile", "p", "default", "AWS profile")
	cmd.PersistentFlags().StringP("cluster", "c", "default", "ECS Cluster Name")
}
