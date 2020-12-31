package node

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (h *handler) GetByID(ctx echo.Context, nodeID uint64) (nodeDTO.GetResponseDTO, error) {
	nodeData, err := h.domain.Node.GetByID(ctx, nodeID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.NotFoundError) {
			statusCode = http.StatusNotFound
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return nodeDTO.GetResponseDTO{}, err
	}

	sensorsData, err := h.domain.NodeSensor.AllByNodeID(ctx, nodeID)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return nodeDTO.GetResponseDTO{}, err
	}

	resp := nodeDTO.GetResponseDTO{
		ResponseDTO: nodeData.ToResponseDTO(),
		Sensors:     sensorsData.ToResponsesDTO(),
	}

	return resp, nil
}
