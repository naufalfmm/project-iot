package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/node/repository"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		CheckToken(ctx echo.Context, token string) (dao.Node, error)
		Create(ctx echo.Context, create nodeDTO.CreateDTO) (dao.Node, error)
		All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (dao.Nodes, error)
		Count(ctx echo.Context) (int64, error)
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
