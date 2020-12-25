package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) BulkInsert(ctx echo.Context, insertedDataList dao.SensorDataList) (dao.SensorDataList, error) {
	err := r.baseInsert(ctx, insertedDataList)
	if err != nil {
		return dao.SensorDataList{}, err
	}

	return insertedDataList, nil
}
