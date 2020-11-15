package node

import (
	"time"

	"github.com/naufalfmm/project-iot/model/dto/sensorGroup"
)

type (
	ResponseDTO struct {
		ID          uint64     `json:"id"`
		Label       string     `json:"label"`
		Location    *string    `json:"location"`
		Token       string     `json:"token"`
		Type        string     `json:"type"`
		GroupNumber uint64     `json:"group_number"`
		CreatedAt   time.Time  `json:"created_at"`
		CreatedBy   uint64     `json:"created_by"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		UpdatedBy   *uint64    `json:"updated_by,omitempty"`
		DeletedAt   *time.Time `json:"deleted_at,omitempty"`
		DeletedBy   *time.Time `json:"deleted_by,omitempty"`
	}

	CreateResponseDTO struct {
		ResponseDTO
		SensorGroups []sensorGroup.ResponseDTO `json:"sensor_groups"`
	}
)
