package dao

import nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"

type NodeSensors []NodeSensor

func (NodeSensors) TableName() string {
	return "node_sensors"
}

func (nss NodeSensors) ToResponsesDTO() []nodeSensorDTO.ResponseDTO {
	resp := make([]nodeSensorDTO.ResponseDTO, len(nss))

	for i := 0; i < len(nss); i++ {
		resp[i] = nss[i].ToResponseDTO()
	}

	return resp
}
