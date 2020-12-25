package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		Insert(ctx echo.Context, insertedData dao.SensorData) (dao.SensorData, error)
		BulkInsert(ctx echo.Context, insertedDataList dao.SensorDataList) (dao.SensorDataList, error)
	}

	repository struct {
		resource resource.Resource
	}
)

func New(resource resource.Resource) (Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}

func (r *repository) baseInsert(ctx echo.Context, data interface{}) error {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return err
	}

	err = orm.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}
