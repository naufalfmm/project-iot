package sensorData

import (
	"time"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

type (
	ResponseDTO struct {
		ID          uint64     `json:"id"`
		PH          float64    `json:"ph"`
		Temperature float64    `json:"temp"`
		TDS         float64    `json:"tds"`
		NodeID      uint64     `json:"node_id"`
		GroupTh     uint64     `json:"group_th"`
		Timestamp   time.Time  `json:"timestamp"`
		CreatedAt   time.Time  `json:"created_at"`
		CreatedBy   uint64     `json:"created_by"`
		UpdatedAt   *time.Time `json:"updated_at"`
		UpdatedBy   *uint64    `json:"updated_by"`
		DeletedAt   *time.Time `json:"deleted_at"`
		DeletedBy   *uint64    `json:"deleted_by"`
	}
)

type (
	PostFromNodeResponseDTO struct {
		Node nodeDTO.ResponseDTO
		Data []ResponseDTO
	}
)
