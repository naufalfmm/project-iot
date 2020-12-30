package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Create(ctx echo.Context, newSensor dao.NodeSensor) (dao.NodeSensor, error) {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.NodeSensor{}, err
	}

	err = orm.Create(&newSensor).Error
	if err != nil {
		return dao.NodeSensor{}, err
	}

	return newSensor, nil
}
