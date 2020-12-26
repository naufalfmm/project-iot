package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) All(ctx echo.Context, params nodeDTO.AllRequestParamsDTO) (dao.Nodes, error) {
	resultDAO, err := s.repository.All(ctx, params)
	if err != nil {
		return nil, err
	}

	return resultDAO, nil
}
