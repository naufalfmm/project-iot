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
		Code       string    `json:"code"`
		Category   string    `json:"category"`
		Value      float64   `json:"value"`
		Unit       string    `json:"unit"`
		GroupLabel string    `json:"group_label"`
		GroupTh    uint32    `json:"group_th"`
		Timestamp  time.Time `json:"timestamp"`
	}
)

type (
	PostFromNodeResponseDTO struct {
		Node nodeDTO.ResponseDTO `json:"node"`
		Data []ResponseDTO       `json:"data"`
	}
)
