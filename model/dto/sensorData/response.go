package sensorData

import (
	"time"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

type (
	ResponseDTO struct {
		ID         uint64    `json:"id,omitempty"`
		NodeID     uint64    `json:"node_id"`
		NodeLabel  string    `json:"node_label"`
		GroupTh    uint64    `json:"group_th"`
		SensorCode string    `json:"sensor_code"`
		SensorType string    `json:"sensor_type"`
		Value      float64   `json:"value"`
		Unit       string    `json:"unit"`
		Timestamp  time.Time `json:"timestamp"`
	}
)

type (
	PostFromNodeResponseDTO struct {
		Node nodeDTO.ResponseDTO
		Data []ResponseDTO
	}
)
