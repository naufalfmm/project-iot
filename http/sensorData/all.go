package sensorData

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/defaultResp"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (c *Controller) All(ctx echo.Context) error {
	params := sensorDataDTO.NewAllRequestParamsDTO(ctx)

	if len(params.Sort) == 1 && params.Sort[0] == "-created_at" {
		params.Sort = []string{"-timestamp"}
	}

	all, err := c.SensorData.All(ctx, params)
	if err != nil {
		return defaultResp.CreateResp(ctx, err)
	}

	return defaultResp.CreateResp(ctx, all)
}
