package node

import (
	"math"
	"time"

	"github.com/naufalfmm/project-iot/common/paging"
	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

type (
	GroupNumberResponseParamDTO struct {
		Min *uint64 `json:"min,omitempty"`
		Max *uint64 `json:"max,omitempty"`
	}

	AllResponseParamsDTO struct {
		paging.PagingRequest
		GroupNumber *GroupNumberResponseParamDTO `json:"group_number,omitempty"`
	}
)

type (
	ResponseDTO struct {
		ID        uint64    `json:"id"`
		Label     string    `json:"label"`
		Location  *string   `json:"location"`
		Token     string    `json:"token"`
		Type      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		CreatedBy string    `json:"created_by"`
	}

	GetResponseDTO struct {
		ResponseDTO
		Sensors []nodeSensorDTO.ResponseDTO `json:"sensors"`
	}
)

type (
	GetAllResponseDTO struct {
		Count       int                  `json:"count"`
		CurrentPage int                  `json:"currentPage"`
		TotalPages  int                  `json:"totalPages"`
		Params      AllResponseParamsDTO `json:"params"`
		Items       []ResponseDTO        `json:"items"`
	}
)

func NewGetAllResponseDTO(params AllRequestParamsDTO, items []ResponseDTO, countTotal int64) GetAllResponseDTO {
	totalPagesInFloat := float64(countTotal) / float64(params.Limit)
	totalPagesInFloat = math.Ceil(totalPagesInFloat)

	totalPages := int(totalPagesInFloat)

	return GetAllResponseDTO{
		Count:       len(items),
		CurrentPage: params.Page,
		TotalPages:  totalPages,
		Params:      params.ToAllResponseParamsDTO(),
		Items:       items,
	}
}
