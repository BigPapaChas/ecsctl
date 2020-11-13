package ecsutils

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/ecsctl/ecsctl/internal/awsutils"
)

var ecsSvc *ecs.ECS

func DescribeContainerInstances(cluster *string) ([]*ecs.ContainerInstance, error) {
	instanceArns, err := listContainerInstances(cluster)
	if err != nil {
		return nil, err
	}

	return describeContainerInstances(instanceArns, cluster)
}

func DescribeECSServices(cluster *string) ([]*ecs.Service, error) {
	serviceArns, err := listServices(cluster)
	if err != nil {
		return nil, err
	}

	return describeServices(serviceArns, cluster)
}

func getECSClient() *ecs.ECS {
	if ecsSvc == nil {
		ecsSvc = ecs.New(awsutils.GetSession())
	}
	return ecsSvc
}

func describeContainerInstances(instanceArns []*string, cluster *string) ([]*ecs.ContainerInstance, error) {
	client := getECSClient()
	var containerInstances []*ecs.ContainerInstance
	for start := 0; start < len(instanceArns); start += 100 {
		end := start + 100
		if end > len(instanceArns) {
			end = len(instanceArns)
		}
		params := &ecs.DescribeContainerInstancesInput{
			Cluster: cluster,
			ContainerInstances: instanceArns[start:end],
		}
		if output, err := client.DescribeContainerInstances(params); err != nil {
			return nil, err
		} else {
			containerInstances = append(containerInstances, output.ContainerInstances...)
		}
	}
	return containerInstances, nil
}

func listContainerInstances(cluster *string) ([]*string, error) {
	client := getECSClient()
	var containerInstanceArns []*string

	params :=  &ecs.ListContainerInstancesInput{Cluster: cluster}
	err := client.ListContainerInstancesPages(params, func(output *ecs.ListContainerInstancesOutput, b bool) bool {
		containerInstanceArns = append(containerInstanceArns, output.ContainerInstanceArns...)
		return true
	})

	return containerInstanceArns, err
}

func describeServices(serviceArns []*string, cluster *string) ([]*ecs.Service, error) {
	client := getECSClient()
	var services []*ecs.Service
	for start := 0; start < len(serviceArns); start += 10 {
		end := start + 10
		if end > len(serviceArns) {
			end = len(serviceArns)
		}
		params := &ecs.DescribeServicesInput{
			Cluster:  cluster,
			Services: serviceArns[start:end],
		}
		if output, err := client.DescribeServices(params); err != nil {
			return nil, err
		} else {
			services = append(services, output.Services...)
		}
	}

	return services, nil
}

func listServices(cluster *string) ([]*string, error) {
	client := getECSClient()
	params := &ecs.ListServicesInput{Cluster: cluster}
	var serviceArns []*string

	err := client.ListServicesPages(params, func(output *ecs.ListServicesOutput, b bool) bool {
		serviceArns = append(serviceArns, output.ServiceArns...)
		return true
	})

	return serviceArns, err
}