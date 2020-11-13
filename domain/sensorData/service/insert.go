package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"

	sensorDataDTO "github.com/naufalfmm/project-iot/model/dto/sensorData"
)

func (s *service) Insert(ctx echo.Context, req sensorDataDTO.CreateDTO) (sensorDataDTO.ResponseDTO, error) {
	sensorData := dao.NewFromCreateDTO(req)

	data, err := s.repository.Insert(ctx, sensorData)
	if err != nil {
		return sensorDataDTO.ResponseDTO{}, err
	}

	return data.ToResponseDTO(), nil
}
