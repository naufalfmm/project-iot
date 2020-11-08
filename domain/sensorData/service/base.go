package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/sensorData/repository"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		Insert(ctx echo.Context, req sensorDataDTO.CreateDTO, doer uint64) (sensorDataDTO.ResponseDTO, error)
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
