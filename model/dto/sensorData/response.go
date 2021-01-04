package sensorData

import (
	"time"

	"github.com/naufalfmm/project-iot/common/paging"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

type (
	TimestampResponseParamDTO struct {
		Min *time.Time `json:"min,omitempty"`
		Max *time.Time `json:"max,omitempty"`
	}

	AllResponseParamsDTO struct {
		paging.PagingRequest
		Timestamp *TimestampResponseParamDTO `json:"timestamp,omitempty"`
		NodeID    *uint64                    `json:"node_id,omitempty"`
	}
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

type (
	GetAllResponseDTO struct {
		Next   bool                 `json:"next"`
		Count  int                  `json:"count"`
		Params AllResponseParamsDTO `json:"params"`
		Items  []ResponseDTO        `json:"items"`
	}
)

func NewGetAllResponseDTO(params AllRequestParamsDTO, items []ResponseDTO, next bool) GetAllResponseDTO {
	return GetAllResponseDTO{
		Next:   next,
		Count:  len(items),
		Params: params.ToAllResponseParamsDTO(),
		Items:  items,
	}
}
