package nodeSensor

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
)

func (h *handler) ToggleActive(ctx echo.Context, sensorID uint64) error {
	err := h.domain.NodeSensor.ToggleActive(ctx, sensorID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.NotFoundError) {
			statusCode = http.StatusBadRequest
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return err
	}

	ctx.Set(consts.ResponseCode, http.StatusNoContent)
	return nil
}
