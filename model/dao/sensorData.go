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
	GroupTh     uint64    `gorm:"not null"`
	Timestamp   time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	CreatedBy   uint64    `gorm:"not null"`
	UpdatedAt   *time.Time
	UpdatedBy   *uint64
	DeletedAt   *time.Time
	DeletedBy   *uint64
}

func (SensorData) TableName() string {
	return "sensor_data"
}

func (s SensorData) ToResponseDTO() sensorData.ResponseDTO {
	return sensorData.ResponseDTO(s)
}

func NewFromCreateDTO(s sensorData.CreateDTO, doer uint64) SensorData {
	now := time.Now()

	return SensorData{
		PH:          s.PH,
		Temperature: s.Temp,
		TDS:         s.TDS,
		NodeID:      s.NodeID,
		GroupTh:     s.GroupTh,
		Timestamp:   now,
		CreatedAt:   now,
		CreatedBy:   doer,
	}
}
