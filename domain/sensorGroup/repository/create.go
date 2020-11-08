package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Create(ctx echo.Context, newGroup dao.SensorGroup) (dao.SensorGroup, error) {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.SensorGroup{}, err
	}

	err = orm.Create(&newGroup).Error
	if err != nil {
		return dao.SensorGroup{}, err
	}

	return newGroup, nil
}
