package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (r *repository) All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (dao.Nodes, error) {
	var results dao.Nodes

	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return nil, err
	}

	offset := (params.Page - 1) * params.Limit

	where, dataInterface := params.GroupNumberWhereQuery()
	if where != "" {
		orm = orm.Where(where, dataInterface)
	}

	sortQuery := params.SortQuery()
	if sortQuery != "" {
		orm = orm.Order(sortQuery)
	}

	err = orm.
		Where("is_deleted = ?", false).
		Limit(params.Limit).
		Offset(offset).
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
