package http

import (
	"github.com/naufalfmm/project-iot/http/node"
	"github.com/naufalfmm/project-iot/http/sensorData"
	"github.com/naufalfmm/project-iot/http/user"
)

type (
	Controllers struct {
		SensorData sensorData.Controller
		Node       node.Controller
		User       user.Controller
	}
)
