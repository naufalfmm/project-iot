package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"
)

func (s *service) Create(ctx echo.Context, create sensorGroupDTO.CreateDTO) (dao.SensorGroup, error) {
	newSensorGroupDTO := dao.NewSensorGroupFromCreateDTO(create)

	newSensorGroup, err := s.repository.Create(ctx, newSensorGroupDTO)
	if err != nil {
		return dao.SensorGroup{}, err
	}

	return newSensorGroup, nil
}
