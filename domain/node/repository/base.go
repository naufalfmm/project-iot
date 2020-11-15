package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		GetByToken(ctx echo.Context, token string) (dao.Node, error)
		Create(ctx echo.Context, newNode dao.Node) (dao.Node, error)
		All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (dao.Nodes, error)
		Count(ctx echo.Context) (int64, error)
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
