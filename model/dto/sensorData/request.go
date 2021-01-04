package sensorData

import (
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/paging"
	"gorm.io/gorm"
)

type (
	AllRequestParamsDTO struct {
		paging.PagingRequest
		NodeID       *uint64
		TimestampMin *time.Time
		TimestampMax *time.Time
	}
)

func NewAllRequestParamsDTO(ctx echo.Context) AllRequestParamsDTO {
	var nodeID *uint64

	pagingParam, _ := paging.NewPagingRequest(ctx)
	tmMin, tmMax := getTimestampFilter(ctx)

	nodeIDQuery := ctx.QueryParam("node_id")
	if nodeIDQuery != "" {
		nodeIDParse, _ := strconv.ParseUint(strings.TrimSpace(nodeIDQuery), 10, 64)
		nodeID = &nodeIDParse
	}

	arp := AllRequestParamsDTO{
		pagingParam,
		nodeID,
		tmMin,
		tmMax,
	}

	return arp
}

func getTimestampFilter(ctx echo.Context) (*time.Time, *time.Time) {
	var (
		tmMin, tmMax *time.Time
	)

	timestampQp := ctx.QueryParam("timestamp")
	if timestampQp == "" {
		return nil, nil
	}

	timestamps := strings.Split(timestampQp, ",")

	timestampMin, err := time.Parse(time.RFC3339, timestamps[0])
	if err == nil {
		tmMin = &timestampMin
	}

	if len(timestamps) == 2 {
		timestampMax, err := time.Parse(time.RFC3339, timestamps[1])
		if err == nil {
			tmMax = &timestampMax
		}
	}

	return tmMin, tmMax
}

func (arp AllRequestParamsDTO) WhereQuery(orm *gorm.DB) *gorm.DB {
	if arp.TimestampMin != nil && arp.TimestampMax != nil {
		orm = orm.Where("timestamp >= ? AND timestamp <= ?", *arp.TimestampMin, *arp.TimestampMax)
	}

	if arp.TimestampMin == nil && arp.TimestampMax != nil {
		orm = orm.Where("timestamp <= ?", *arp.TimestampMax)
	}

	if arp.TimestampMax == nil && arp.TimestampMin != nil {
		orm = orm.Where("timestamp >= ?", *arp.TimestampMin)
	}

	if arp.NodeID != nil {
		orm = orm.Where("node_id = ?", *arp.NodeID)
	}

	return orm
}

func (arp AllRequestParamsDTO) ToAllResponseParamsDTO() AllResponseParamsDTO {
	var timestamp *TimestampResponseParamDTO

	pagingReq := paging.PagingRequest{
		Page:  arp.Page,
		Limit: arp.Limit,
		Sort:  arp.Sort,
	}

	if arp.TimestampMin != nil || arp.TimestampMax != nil {
		timestamp = &TimestampResponseParamDTO{
			Min: arp.TimestampMin,
			Max: arp.TimestampMax,
		}
	}

	return AllResponseParamsDTO{
		pagingReq,
		timestamp,
		arp.NodeID,
	}
}

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
		Token     string          `validate:"required"`
		Data      map[int]float64 `validate:"required"`
		Timestamp time.Time       `validate:"required"`
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
