package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		GetByToken(ctx echo.Context, token string) (dao.Node, error)
		Create(ctx echo.Context, newNode dao.Node) (dao.Node, error)
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
