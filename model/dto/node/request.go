package node

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/login"
	"github.com/naufalfmm/project-iot/common/paging"
)

type (
	AllRequestParamsDTO struct {
		paging.PagingRequest
		GroupNumberMin *uint64
		GroupNumberMax *uint64
	}
)

func NewAllRequestParamsDTO(ctx echo.Context) AllRequestParamsDTO {
	pagingParam, _ := paging.NewPagingRequest(ctx)
	gnMin, gnMax := getGroupNumberFilter(ctx)

	arp := AllRequestParamsDTO{
		pagingParam,
		gnMin,
		gnMax,
	}

	return arp
}

func getGroupNumberFilter(ctx echo.Context) (*uint64, *uint64) {
	var (
		gnMax, gnMin *uint64
	)

	groupNumberQp := ctx.QueryParam("group_number")
	groupNumbers := strings.Split(groupNumberQp, ",")

	if groupNumberQp == "" {
		return nil, nil
	}

	groupNumberMin, _ := strconv.ParseUint(strings.TrimSpace(groupNumbers[0]), 10, 64)
	gnMin = &groupNumberMin

	if len(groupNumbers) == 2 {
		groupNumberMax, _ := strconv.ParseUint(strings.TrimSpace(groupNumbers[1]), 10, 64)
		gnMax = &groupNumberMax
	}

	return gnMin, gnMax
}

func (arp AllRequestParamsDTO) GroupNumberWhereQuery() (string, []uint64) {
	if arp.GroupNumberMin != nil && arp.GroupNumberMax != nil {
		return "group_number >= ? AND group_number <= ?", []uint64{*arp.GroupNumberMin, *arp.GroupNumberMax}
	}

	if arp.GroupNumberMin == nil && arp.GroupNumberMax != nil {
		return "group_number <= ?", []uint64{*arp.GroupNumberMax}
	}

	if arp.GroupNumberMax == nil && arp.GroupNumberMin != nil {
		return "group_number >= ?", []uint64{*arp.GroupNumberMin}
	}

	return "", []uint64{}
}

func (arp AllRequestParamsDTO) ToAllResponseParamsDTO() AllResponseParamsDTO {
	var groupNumber *GroupNumberResponseParamDTO

	pagingReq := paging.PagingRequest{
		Page:  arp.Page,
		Limit: arp.Limit,
		Sort:  arp.Sort,
	}

	if arp.GroupNumberMin != nil || arp.GroupNumberMax != nil {
		groupNumber = &GroupNumberResponseParamDTO{
			Min: arp.GroupNumberMin,
			Max: arp.GroupNumberMax,
		}
	}

	return AllResponseParamsDTO{
		pagingReq,
		groupNumber,
	}
}

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
		Label             string   `json:"label" validate:"required"`
		Location          *string  `json:"location"`
		Type              string   `json:"type" validate:"required"`
		SensorGroupLabels []string `json:"sensor_group_labels" validate:"required"`
	}
)

type (
	CreateRequestDTO struct {
		Body CreateRequestBodyDTO
		By   login.ClientJWTDTO
	}
)
