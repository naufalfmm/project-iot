package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Insert(ctx echo.Context, insertedData dao.SensorData) (dao.SensorData, error) {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.SensorData{}, err
	}

	err = orm.Create(&insertedData).Error
	if err != nil {
		return dao.SensorData{}, err
	}

	return insertedData, nil
}
