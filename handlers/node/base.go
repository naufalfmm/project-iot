package node

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Handler interface {
		CreateComplete(ctx echo.Context, createReq nodeDTO.CreateRequestDTO) (nodeDTO.CreateResponseDTO, error)
		All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (nodeDTO.GetAllResponseDTO, error)
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
