package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Repository interface {
		AllByNodeID(ctx echo.Context, nodeID uint64) (dao.NodeSensors, error)
		Create(ctx echo.Context, newSensor dao.NodeSensor) (dao.NodeSensor, error)
		ToggleActive(ctx echo.Context, sensorID uint64) (int64, error)
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
