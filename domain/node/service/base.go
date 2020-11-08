package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/node/repository"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		CheckToken(ctx echo.Context, token string) (nodeDTO.ResponseDTO, error)
		Create(ctx echo.Context, create nodeDTO.CreateDTO) (nodeDTO.ResponseDTO, error)
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
