package nodeSensor

import (
	"net/http"
	"strconv"

	"github.com/naufalfmm/project-iot/common/defaultResp"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ToggleActive(ctx echo.Context) error {
	sensorID, err := strconv.ParseUint(ctx.Param("sensorId"), 10, 64)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return defaultResp.CreateResp(ctx, nil)
	}

	err = c.NodeSensor.ToggleActive(ctx, sensorID)
	if err != nil {
		return defaultResp.CreateResp(ctx, err)
	}

	return defaultResp.CreateResp(ctx, nil)
}
