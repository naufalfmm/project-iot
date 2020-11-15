package service

import (
	"github.com/labstack/echo/v4"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) ([]nodeDTO.ResponseDTO, error) {
	resultDAO, err := s.repository.All(ctx, params)
	if err != nil {
		return nil, err
	}

	return resultDAO.ToResponsesDTO(), nil
}
