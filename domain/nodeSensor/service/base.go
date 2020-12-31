package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/nodeSensor/repository"
	"github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		AllByNodeID(ctx echo.Context, nodeID uint64) (dao.NodeSensors, error)
		Create(ctx echo.Context, create nodeSensorDTO.CreateDTO) (dao.NodeSensor, error)
		ToggleActive(ctx echo.Context, sensorID uint64) error
	}

	service struct {
		resource   resource.Resource
		repository repository.Repository
		subdomain  subdomain.Subdomain
	}
)

func New(resource resource.Resource, repository repository.Repository, subdomain subdomain.Subdomain) (Service, error) {
	return &service{
		resource:   resource,
		repository: repository,
		subdomain:  subdomain,
	}, nil
}
