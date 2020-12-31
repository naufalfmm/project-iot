package nodeSensor

import (
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"

	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

func (h *handler) Create(ctx echo.Context, req nodeSensorDTO.CreateRequestDTO) (nodeSensorDTO.ResponseDTO, error) {
	newSensorDTO := req.ToCreateDTO()

	newSensor, err := h.domain.NodeSensor.Create(ctx, newSensorDTO)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return nodeSensorDTO.ResponseDTO{}, err
	}

	return newSensor.ToResponseDTO(), err
}
