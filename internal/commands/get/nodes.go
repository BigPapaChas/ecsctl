package get

import (
	"github.com/ecsctl/ecsctl/internal/awsutils/ecsutils"
	"github.com/ecsctl/ecsctl/internal/configutils"
	"github.com/ecsctl/ecsctl/internal/printerutils"
	"github.com/spf13/cobra"
)

type NodeOptions struct {
	Nodes   []string
	Cluster string
}



func newCmdGetNodes() *cobra.Command {
	o := &NodeOptions{}
	cmd := &cobra.Command{
		Use:   "nodes",
		Aliases: []string{"node", "instances"},
		Short: "Retrieves node information of an ECS cluster",
		Long:  `Displays information of the container instances of the ECS cluster`,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.loadFlags(cmd)
			o.loadArgs()
			if err := o.validateOptions(); err != nil {
				return err
			}
			return o.run()
		},
	}
	return cmd
}

func (o *NodeOptions) loadFlags(cmd *cobra.Command) {
	o.Cluster, _ = configutils.GetCluster(cmd)
}

func (o *NodeOptions) loadArgs() {

}

func (o *NodeOptions) validateOptions() error {
	return nil
}

func (o *NodeOptions) run() error {
	return getNodes(o)
}


func getNodes(o *NodeOptions) error {
	nodes, err := ecsutils.DescribeContainerInstances(&o.Cluster)
	if err != nil {
		return err
	}
	printerutils.PrintNodes(nodes)
	return nil
}
