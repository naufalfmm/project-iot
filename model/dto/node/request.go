package node

import (
	"github.com/naufalfmm/project-iot/common/login"
)

type (
	CreateDTO struct {
		Label       string
		Location    *string
		Type        string
		GroupNumber uint64
		By          uint64
	}
)

type (
	CreateRequestBodyDTO struct {
		Label       string   `json:"label" validate:"required"`
		Location    *string  `json:"location"`
		Type        string   `json:"type" validate:"required"`
		SensorGroup []string `json:"sensor_group_labels" validate:"required"`
	}
)

type (
	CreateRequestDTO struct {
		Body CreateRequestBodyDTO
		By   login.ClientJWTDTO
	}
)
