package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/sensorGroup/repository"
	"github.com/naufalfmm/project-iot/model/dao"
	sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		Create(ctx echo.Context, create sensorGroupDTO.CreateDTO) (dao.SensorGroup, error)
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
