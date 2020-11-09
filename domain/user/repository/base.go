package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		Create(ctx echo.Context, newUser dao.User) (dao.User, error)
		Get(ctx echo.Context, whereQuery dao.User) (dao.User, error)
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
