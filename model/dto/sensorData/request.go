package sensorData

type (
	CreateDTO struct {
		PH      float64
		TDS     float64
		Temp    float64
		GroupTh uint64
		NodeID  uint64
	}
)

type (
	PostDTO struct {
		PH   float64 `validate:"required"`
		TDS  float64 `validate:"required"`
		Temp float64 `validate:"required"`
	}

	PostFromNodeRequestDTO struct {
		Token string    `validate:"required"`
		Data  []PostDTO `validate:"dive"`
	}
)

func (p PostDTO) ToCreateDTO(groupTh uint64, nodeID uint64) CreateDTO {
	return CreateDTO{
		PH:      p.PH,
		TDS:     p.TDS,
		Temp:    p.Temp,
		GroupTh: groupTh,
		NodeID:  nodeID,
	}
}
