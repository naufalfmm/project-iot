package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/sensorData/repository"
	"github.com/naufalfmm/project-iot/model/dao"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		Insert(ctx echo.Context, req sensorDataDTO.CreateDTO) (dao.SensorData, error)
		BulkInsert(ctx echo.Context, reqs []sensorDataDTO.CreateDTO) (dao.SensorDataList, error)
		AllNext(ctx echo.Context, params sensorDataDTO.AllRequestParamsDTO) (bool, dao.SensorDataList, error)
	}

	service struct {
		resource   resource.Resource
		repository repository.Repository
	}
)

func New(resource resource.Resource, repository repository.Repository) (Service, error) {
	return &service{
		resource:   resource,
		repository: repository,
	}, nil
}
