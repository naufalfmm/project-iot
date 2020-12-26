package node

import (
	"math"
	"time"

	"github.com/naufalfmm/project-iot/common/paging"

	"github.com/naufalfmm/project-iot/model/dto/sensorGroup"
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
		ID          uint64     `json:"id"`
		Label       string     `json:"label"`
		Location    *string    `json:"location"`
		Token       string     `json:"token"`
		Type        string     `json:"type"`
		GroupNumber uint64     `json:"group_number"`
		CreatedAt   time.Time  `json:"created_at"`
		CreatedBy   string     `json:"created_by"`
		UpdatedAt   *time.Time `json:"updated_at,omitempty"`
		UpdatedBy   *string    `json:"updated_by,omitempty"`
		DeletedAt   *time.Time `json:"deleted_at,omitempty"`
		DeletedBy   *string    `json:"deleted_by,omitempty"`
		IsDeleted   bool       `json:"is_deleted"`
	}

	CreateResponseDTO struct {
		ResponseDTO
		SensorGroups []sensorGroup.ResponseDTO `json:"sensor_groups"`
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
