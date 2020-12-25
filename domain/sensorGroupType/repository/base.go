package repository

import (
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface{
		AllByNodeID(ctx echo.Context, nodeID uint64) (dao.SensorGroupType, error)
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
