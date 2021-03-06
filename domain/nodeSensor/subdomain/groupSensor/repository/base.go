package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		GetUniqueByNodeID(ctx echo.Context, nodeID uint64) (dao.GroupSensors, error)
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
