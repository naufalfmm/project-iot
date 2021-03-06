package http

import (
	"github.com/naufalfmm/project-iot/http/auth"
	"github.com/naufalfmm/project-iot/http/node"
	"github.com/naufalfmm/project-iot/http/nodeSensor"
	"github.com/naufalfmm/project-iot/http/sensorData"
	"github.com/naufalfmm/project-iot/http/user"
)

type (
	Controllers struct {
		SensorData sensorData.Controller
		Node       node.Controller
		User       user.Controller
		NodeSensor nodeSensor.Controller
		Auth       auth.Controller
	}
)
