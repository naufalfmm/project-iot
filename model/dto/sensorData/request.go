package sensorData

import "time"

type (
	CreateDTO struct {
		NodeID     uint64
		NodeLabel  string
		Code       string
		Category   string
		Value      float64
		Unit       string
		GroupLabel string
		GroupTh    uint32
		Timestamp  time.Time
		CreatedBy  string
	}
)

type (
	PostFromNodeRequestDTO struct {
		Token     string    `validate:"required"`
		Data      []float64 `validate:"required"`
		Timestamp time.Time `validate:"required"`
	}
)

func (c CreateDTO) ToResponseDTO() ResponseDTO {
	return ResponseDTO{
		NodeID:     c.NodeID,
		NodeLabel:  c.NodeLabel,
		Code:       c.Code,
		Category:   c.Category,
		Value:      c.Value,
		Unit:       c.Unit,
		GroupLabel: c.GroupLabel,
		GroupTh:    c.GroupTh,
		Timestamp:  c.Timestamp,
	}
}
