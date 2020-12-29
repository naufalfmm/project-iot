package node

import (
	"github.com/labstack/echo/v4"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (h *handler) All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (nodeDTO.GetAllResponseDTO, error) {
	all, err := h.domain.Node.All(ctx, params)
	if err != nil {
		return nodeDTO.GetAllResponseDTO{}, err
	}

	count, err := h.domain.Node.Count(ctx)
	if err != nil {
		return nodeDTO.GetAllResponseDTO{}, err
	}

	allDTO := all.ToResponsesDTO()

	resp := nodeDTO.NewGetAllResponseDTO(params, allDTO, count)

	return resp, nil
}
