package nodeSensor

import (
	"github.com/naufalfmm/project-iot/handlers/nodeSensor"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		NodeSensor nodeSensor.Handler
		Resource   resource.Resource
	}
)
