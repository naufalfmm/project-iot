package handlers

import (
	"github.com/naufalfmm/project-iot/handlers/node"
	"github.com/naufalfmm/project-iot/handlers/sensorData"
	"github.com/naufalfmm/project-iot/handlers/user"
)

type (
	Handlers struct {
		SensorData sensorData.Handler
		Node       node.Handler
		User       user.Handler
	}
)
