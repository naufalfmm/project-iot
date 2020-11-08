package sensorData

import (
	"github.com/naufalfmm/project-iot/handlers/sensorData"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		SensorData sensorData.Handler
		Resource   resource.Resource
	}
)
