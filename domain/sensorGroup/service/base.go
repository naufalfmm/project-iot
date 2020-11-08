package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/sensorGroup/repository"
	sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		Create(ctx echo.Context, create sensorGroupDTO.CreateDTO) (sensorGroupDTO.ResponseDTO, error)
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
