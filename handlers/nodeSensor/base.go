package nodeSensor

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain"
	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Handler interface {
		Create(ctx echo.Context, req nodeSensorDTO.CreateRequestDTO) (nodeSensorDTO.ResponseDTO, error)
		ToggleActive(ctx echo.Context, sensorID uint64) error
	}

	handler struct {
		domain   domain.Domain
		resource resource.Resource
	}
)

func NewHandler(domain domain.Domain, resource resource.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
