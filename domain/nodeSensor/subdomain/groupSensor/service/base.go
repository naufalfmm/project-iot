package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain/groupSensor/repository"
	"github.com/naufalfmm/project-iot/model/dao"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		GetUniqueByNodeID(ctx echo.Context, nodeID uint64) (dao.GroupSensors, error)
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
