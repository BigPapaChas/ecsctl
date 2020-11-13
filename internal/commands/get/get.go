package get

import (
	"github.com/ecsctl/ecsctl/internal/awsutils"
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves resources of an ECS cluster",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var region, profile string
			var err error
			if region, err = configutils.GetRegion(cmd); err != nil {
				return err
			}
			if profile, err = configutils.GetProfile(cmd); err != nil {
				return err
			}
			return awsutils.InitAWSSession(region, profile)
		},
	}

	// Add the common persistent flags
	configutils.AddCommonPersistentFlags(cmd)

	cmd.AddCommand(newCmdGetNodes())
	cmd.AddCommand(newCmdGetServices())
	return cmd
}
