package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (s *service) BulkInsert(ctx echo.Context, reqs []sensorDataDTO.CreateDTO) (dao.SensorDataList, error) {
	sensorDataList := dao.NewFromCreatesDTO(reqs)

	data, err := s.repository.BulkInsert(ctx, sensorDataList)
	if err != nil {
		return dao.SensorDataList{}, err
	}

	return data, nil
}
