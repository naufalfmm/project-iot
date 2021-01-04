package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (s *service) AllNext(ctx echo.Context, params sensorDataDTO.AllRequestParamsDTO) (bool, dao.SensorDataList, error) {
	return s.repository.AllNext(ctx, params)
}
