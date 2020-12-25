package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Insert(ctx echo.Context, insertedData dao.SensorData) (dao.SensorData, error) {
	err := r.baseInsert(ctx, insertedData)
	if err != nil {
		return dao.SensorData{}, err
	}

	return insertedData, nil
}
