package nodeSensor

import (
	"errors"
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"

	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

func (h *handler) Create(ctx echo.Context, req nodeSensorDTO.CreateRequestDTO) (nodeSensorDTO.ResponseDTO, error) {
	newSensorDTO := req.ToCreateDTO()

	newSensor, err := h.domain.NodeSensor.Create(ctx, newSensorDTO)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.GroupLabelNotFoundErr) {
			statusCode = http.StatusBadRequest
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return nodeSensorDTO.ResponseDTO{}, err
	}

	return newSensor.ToResponseDTO(), err
}
