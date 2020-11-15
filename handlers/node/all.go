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

	resp := nodeDTO.NewGetAllResponseDTO(params, all, count)

	return resp, nil
}
