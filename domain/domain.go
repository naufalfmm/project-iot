package domain

import (
	node "github.com/naufalfmm/project-iot/domain/node/service"
	sensorData "github.com/naufalfmm/project-iot/domain/sensorData/service"
	sensorGroup "github.com/naufalfmm/project-iot/domain/sensorGroup/service"
	user "github.com/naufalfmm/project-iot/domain/user/service"
)

type (
	Domain struct {
		Node        node.Service
		SensorData  sensorData.Service
		SensorGroup sensorGroup.Service
		User        user.Service
	}
)
