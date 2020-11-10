package dao

import sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"

type SensorGroup struct {
	ID     uint64 `gorm:"PRIMARY_KEY"`
	Label  string `gorm:"not null"`
	NodeID uint64 `gorm:"not null"`
	Th     uint64 `gorm:"not null"`
}

func (SensorGroup) TableName() string {
	return "sensor_groups"
}

func (sg SensorGroup) ToResponseDTO() sensorGroupDTO.ResponseDTO {
	return sensorGroupDTO.ResponseDTO(sg)
}

func NewSensorGroupFromCreateDTO(c sensorGroupDTO.CreateDTO) SensorGroup {
	return SensorGroup{
		Label:  c.Label,
		NodeID: c.NodeID,
		Th:     c.Th,
	}
}