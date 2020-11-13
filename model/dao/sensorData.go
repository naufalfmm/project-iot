package dao

import (
	"time"

	"github.com/naufalfmm/project-iot/model/dto/sensorData"
)

type SensorData struct {
	ID          uint64    `gorm:"PRIMARY_KEY"`
	PH          float64   `gorm:"not null"`
	Temperature float64   `gorm:"not null"`
	TDS         float64   `gorm:"not null"`
	NodeID      uint64    `gorm:"not null"`
	NodeLabel   string    `gorm:"not null"`
	GroupTh     uint64    `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	CreatedBy   uint64    `gorm:"not null"`
	CreatedFrom string    `gorm:"not null"`
	UpdatedAt   *time.Time
}

func (SensorData) TableName() string {
	return "sensor_data"
}

func (s SensorData) ToResponseDTO() sensorData.ResponseDTO {
	return sensorData.ResponseDTO(s)
}

func NewFromCreateDTO(s sensorData.CreateDTO) SensorData {
	now := time.Now()

	return SensorData{
		PH:          s.PH,
		Temperature: s.Temp,
		TDS:         s.TDS,
		NodeID:      s.NodeID,
		NodeLabel:   s.NodeLabel,
		GroupTh:     s.GroupTh,
		Timestamp:   s.Timestamp,
		CreatedAt:   now,
		CreatedBy:   s.CreatedBy,
		CreatedFrom: s.CreatedFrom,
	}
}
