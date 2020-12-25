package dao

import (
	"time"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

type SensorData struct {
	ID         uint64    `gorm:"PRIMARY_KEY"`
	NodeID     uint64    `gorm:"not null"`
	NodeLabel  string    `gorm:"not null"`
	GroupTh    uint64    `gorm:"not null"`
	SensorCode string    `gorm:"not null"`
	SensorType string    `gorm:"not null"`
	Value      float64   `gorm:"not null"`
	Unit       string    `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null"`
	BaseModel
}

func (SensorData) TableName() string {
	return "sensor_data"
}

// func (s SensorData) ToResponseDTO() sensorData.ResponseDTO {
// 	return sensorData.ResponseDTO(s)
// }

func NewFromCreateDTO(s sensorDataDTO.CreateDTO) SensorData {
	now := time.Now()

	return SensorData{
		NodeID:     s.NodeID,
		NodeLabel:  s.NodeLabel,
		GroupTh:    s.GroupTh,
		SensorCode: s.SensorCode,
		SensorType: s.SensorType,
		Value:      s.Value,
		Unit:       s.Unit,
		Timestamp:  s.Timestamp,
		BaseModel: BaseModel{
			CreatedAt: now,
			CreatedBy: s.CreatedBy,
		},
	}
}
