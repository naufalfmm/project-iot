package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	sensorGroupDTO "github.com/naufalfmm/project-iot/model/dto/sensorGroup"
)

func (s *service) Create(ctx echo.Context, create sensorGroupDTO.CreateDTO) (sensorGroupDTO.ResponseDTO, error) {
	newSensorGroup := dao.NewSensorGroupFromCreateDTO(create)

	newSensorGroup, err := s.repository.Create(ctx, newSensorGroup)
	if err != nil {
		return sensorGroupDTO.ResponseDTO{}, err
	}

	return newSensorGroup.ToResponseDTO(), nil
}
