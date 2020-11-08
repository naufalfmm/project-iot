package sensorData

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Handler interface {
		PostFromNode(e echo.Context, req sensorDataDTO.PostFromNodeRequestDTO) (sensorDataDTO.PostFromNodeResponseDTO, error)
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
