package printerutils

import (
	"fmt"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/landoop/tableprinter"
)

var printer *tableprinter.Printer

func init() {
	// Initialize tableprinter
	printer = tableprinter.New(os.Stdout)
	printer.HeaderLine = false
}

func PrintNodes(containerInstances []*ecs.ContainerInstance) {
	var nodes []*node
	for _, inst := range containerInstances {
		nodes = append(nodes, convertToNode(inst))
	}

	sort.Slice(nodes, func(i, j int) bool {
		return *nodes[j].ContainerInstance > *nodes[i].ContainerInstance
	})
	printer.Print(nodes)
}

func PrintServices(ecsServices []*ecs.Service, filters []string) {
	filtersEnabled := filters != nil && len(filters) >= 1
	var services []*service
	for _, svc := range ecsServices {
		if filtersEnabled {
			for _, f := range filters {
				if *svc.ServiceName == f {
					services = append(services, convertToService(svc))
				}
			}
		} else {
			services = append(services, convertToService(svc))
		}
	}

	sort.Slice(services, func(i, j int) bool {
		return *services[j].Name > *services[i].Name
	})
	printer.Print(services)

	if filtersEnabled {
		for _, f := range filters {
			found := false
			for _, s := range services {
				if *s.Name == f {
					found = true
				}
			}
			if !found {
				fmt.Printf("Error: service \"%s\" not found\n", f)
			}
		}
	}
}

func PrintObjects(objects interface{}) {
	printer.Print(objects)
}
