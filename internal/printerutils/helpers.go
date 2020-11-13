package printerutils

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/ecs"
)

// Returns a pointer to shortened duration string
func getDurationString(d time.Duration) *string {
	var duration string

	switch {
	case d.Hours() > 24:
		duration = fmt.Sprintf("%dd", int(d.Hours()/24))
	case d.Hours() >= 1:
		duration = fmt.Sprintf("%dh", int(d.Hours()))
	case d.Minutes() >= 1:
		duration = fmt.Sprintf("%dm", int(d.Minutes()))
	default:
		duration = fmt.Sprintf("%ds", int(d.Seconds()))
	}

	return &duration
}

// Converts a ecs.Service to a service struct
func convertToService(svc *ecs.Service) *service {
	return &service{
		Name:    svc.ServiceName,
		Running: getRunningFraction(svc),
		Type:    svc.LaunchType,
		Status:  svc.Status,
		Created: getDurationString(time.Now().Sub(*svc.CreatedAt)),
	}
}

func getRunningFraction(svc *ecs.Service) *string {
	running := fmt.Sprintf("%d/%d", *svc.RunningCount, *svc.DesiredCount)
	return &running
}

// Converts a ecs.ContainerInstance to a node struct
func convertToNode(instance *ecs.ContainerInstance) *node {
	return &node{
		ContainerInstance: parseContainerInstance(instance),
		EC2Instance:       instance.Ec2InstanceId,
		InstanceType:      getECSAttribute(instance, "ecs.instance-type"),
		AvailabilityZone:  getECSAttribute(instance, "ecs.availability-zone"),
		IP:                getPrivateIp(instance),
		Status:            instance.Status,
		MemoryReservation: getMemoryReservation(instance),
		CPUReservation:    getCPUReservation(instance),
		RunningTasks:      instance.RunningTasksCount,
		Age:               getDurationString(time.Now().Sub(*instance.RegisteredAt)),
		AgentConnected:    instance.AgentConnected,
		Version:           instance.VersionInfo.AgentVersion,
	}
}

func parseContainerInstance(instance *ecs.ContainerInstance) *string {
	arn := *instance.ContainerInstanceArn
	containerInstanceId := arn[strings.LastIndex(arn, "/") + 1:]
	return &containerInstanceId
}

func getMemoryReservation(instance *ecs.ContainerInstance) *string {
	return getResourceReservation(instance, "MEMORY")
}

func getCPUReservation(instance *ecs.ContainerInstance) *string {
	return getResourceReservation(instance, "CPU")
}

func getResourceReservation(instance *ecs.ContainerInstance, resourceType string) *string {
	var reg, rem float64

	for _, resource := range instance.RegisteredResources {
		if *resource.Name == resourceType {
			reg = float64(*resource.IntegerValue)
		}
	}

	for _, resource := range instance.RemainingResources {
		if *resource.Name == resourceType {
			rem = float64(*resource.IntegerValue)
		}
	}

	resPercent := fmt.Sprintf("%d%%", int(((reg-rem)/reg)*100))
	return &resPercent
}

func getPrivateIp(instance *ecs.ContainerInstance) *string {
	for _, attachment := range instance.Attachments {
		for _, detail := range attachment.Details {
			if *detail.Name == "privateIPv4Address" {
				return detail.Value
			}
		}
	}
	return nil
}

func getECSAttribute(instance *ecs.ContainerInstance, key string) *string {
	for _, attr := range instance.Attributes {
		if *attr.Name == key {
			return attr.Value
		}
	}
	return nil
}