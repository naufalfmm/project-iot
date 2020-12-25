package domain

import (
	node "github.com/naufalfmm/project-iot/domain/node/service"
	sensorData "github.com/naufalfmm/project-iot/domain/sensorData/service"
	sensorGroup "github.com/naufalfmm/project-iot/domain/sensorGroup/service"
	sensorGroupType "github.com/naufalfmm/project-iot/domain/sensorGroupType/service"
	user "github.com/naufalfmm/project-iot/domain/user/service"
)

type (
	Domain struct {
		Node            node.Service
		SensorData      sensorData.Service
		SensorGroup     sensorGroup.Service
		SensorGroupType sensorGroupType.Service
		User            user.Service
	}
)
