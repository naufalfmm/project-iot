package http

import (
	"github.com/naufalfmm/project-iot/http/node"
	"github.com/naufalfmm/project-iot/http/sensorData"
)

type (
	Controllers struct {
		SensorData sensorData.Controller
		Node       node.Controller
	}
)
