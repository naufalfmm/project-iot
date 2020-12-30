package sensorData

import "time"

type (
	CreateDTO struct {
		NodeID     uint64
		NodeLabel  string
		GroupTh    uint32
		SensorCode string
		SensorType string
		Value      float64
		Unit       string
		Timestamp  time.Time
		CreatedBy  string
	}
)

type (
	PostFromNodeRequestDTO struct {
		Token     string          `validate:"required"`
		Data      map[int]float64 `validate:"required"`
		Timestamp time.Time       `validate:"required"`
	}
)

func (c CreateDTO) ToResponseDTO() ResponseDTO {
	return ResponseDTO{
		NodeID:     c.NodeID,
		NodeLabel:  c.NodeLabel,
		GroupTh:    c.GroupTh,
		SensorCode: c.SensorCode,
		SensorType: c.SensorType,
		Value:      c.Value,
		Unit:       c.Unit,
		Timestamp:  c.Timestamp,
	}
}

// func (p PostDTO) ToCreateDTO(groupTh uint64, nodeID uint64, nodeLabel string) CreateDTO {
// 	return CreateDTO{
// 		PH:          p.PH,
// 		TDS:         p.TDS,
// 		Temp:        p.Temp,
// 		GroupTh:     groupTh,
// 		NodeID:      nodeID,
// 		NodeLabel:   nodeLabel,
// 		Timestamp:   p.Timestamp,
// 		CreatedBy:   nodeID,
// 		CreatedFrom: "NODE",
// 	}
// }
