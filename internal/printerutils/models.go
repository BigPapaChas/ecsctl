package printerutils

type node struct {
	ContainerInstance *string `header:"container instance"`
	EC2Instance *string `header:"ec2 instance"`
	InstanceType *string `header:"instance type"`
	AvailabilityZone *string `header:"availability zone"`
	IP *string //`header:"ip"`
	Status *string `header:"status"`
	MemoryReservation *string `header:"mem res %"`
	CPUReservation *string `header:"cpu res %"`
	RunningTasks *int64 `header:"running tasks"`
	Age *string `header:"age"`
	AgentConnected *bool `header:"agent connected"`
	Version *string `header:"version"`
}

type service struct {
	Name *string `header:"name"`
	Running *string `header:"running"`
	Type *string `header:"type"`
	Status *string `header:"status"`
	Created *string `header:"created"`
}
