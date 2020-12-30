package domain

import (
	node "github.com/naufalfmm/project-iot/domain/node/service"
	nodeSensor "github.com/naufalfmm/project-iot/domain/nodeSensor/service"
	sensorData "github.com/naufalfmm/project-iot/domain/sensorData/service"
	user "github.com/naufalfmm/project-iot/domain/user/service"
)

type (
	Domain struct {
		Node       node.Service
		SensorData sensorData.Service
		NodeSensor nodeSensor.Service
		User       user.Service
	}
)
