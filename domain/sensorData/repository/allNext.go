package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (r *repository) AllNext(ctx echo.Context, params sensorDataDTO.AllRequestParamsDTO) (bool, dao.SensorDataList, error) {
	var (
		results dao.SensorDataList
		next    bool
	)

	orm := r.resource.DB

	limit := params.Limit + 1
	offset := (params.Page - 1) * limit

	orm = params.WhereQuery(orm)

	sortQuery := params.SortQuery()
	if sortQuery != "" {
		orm = orm.Order(sortQuery)
	}

	err := orm.
		Order("timestamp DESC").
		Limit(limit).
		Offset(offset).
		Find(&results).Error
	if err != nil {
		return false, nil, err
	}

	next = len(results) > params.Limit

	if len(results) > 0 {
		results = results[:len(results)-1]
	}

	return next, results, nil
}
