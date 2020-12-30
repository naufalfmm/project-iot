package dao

import (
	"time"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

type SensorData struct {
	ID         uint64    `gorm:"PRIMARY_KEY"`
	NodeID     uint64    `gorm:"not null"`
	NodeLabel  string    `gorm:"not null"`
	Code       string    `gorm:"not null"`
	Category   string    `gorm:"not null"`
	Value      float64   `gorm:"not null"`
	Unit       string    `gorm:"not null"`
	GroupLabel string    `gorm:"not null"`
	GroupTh    uint32    `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null"`
	BaseModel
}

func (SensorData) TableName() string {
	return "sensor_data"
}

func (s SensorData) ToResponseDTO() sensorDataDTO.ResponseDTO {
	return sensorDataDTO.ResponseDTO{
		ID:         s.ID,
		NodeID:     s.NodeID,
		NodeLabel:  s.NodeLabel,
		Code:       s.Code,
		Category:   s.Category,
		Value:      s.Value,
		Unit:       s.Unit,
		GroupTh:    s.GroupTh,
		GroupLabel: s.GroupLabel,
		Timestamp:  s.Timestamp,
	}
}

func NewFromCreateDTO(s sensorDataDTO.CreateDTO) SensorData {
	now := time.Now()

	return SensorData{
		NodeID:     s.NodeID,
		NodeLabel:  s.NodeLabel,
		Code:       s.Code,
		Category:   s.Category,
		Value:      s.Value,
		Unit:       s.Unit,
		GroupTh:    s.GroupTh,
		GroupLabel: s.GroupLabel,
		Timestamp:  s.Timestamp,
		BaseModel: BaseModel{
			CreatedAt: now,
			CreatedBy: s.CreatedBy,
		},
	}
}
