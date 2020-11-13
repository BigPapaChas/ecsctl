package get

import (
	"github.com/ecsctl/ecsctl/internal/awsutils/ecsutils"
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/ecsctl/ecsctl/internal/printerutils"
	"github.com/spf13/cobra"
)

type ServiceOptions struct {
	Filters []string
	Cluster string
}

func newCmdGetServices() *cobra.Command {
	o := &ServiceOptions{}
	cmd := &cobra.Command{
		Use:   "services",
		Aliases: []string{"service", "svc"},
		Short: "Retrieves ECS service information",
		Long:  `Displays information of the deployed services of the ECS cluster`,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.loadFlags(cmd)
			o.loadArgs(args)
			if err := o.validateOptions(); err != nil {
				return err
			}
			return o.run()
		},
	}
	return cmd
}

func (o *ServiceOptions) loadFlags(cmd *cobra.Command) {
	o.Cluster, _ = configutils.GetCluster(cmd)
}

func (o *ServiceOptions) loadArgs(args []string) {
	o.Filters = args
}

func (o *ServiceOptions) validateOptions() error {
	return nil
}

func (o *ServiceOptions) run() error {
	return getServices(o)
}

func getServices(o *ServiceOptions) error {
	services, err := ecsutils.DescribeECSServices(&o.Cluster)
	if err != nil {
		return err
	}
	printerutils.PrintServices(services, o.Filters)
	return nil
}
