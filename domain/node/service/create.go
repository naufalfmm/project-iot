package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) Create(ctx echo.Context, create nodeDTO.CreateDTO) (nodeDTO.ResponseDTO, error) {
	newNode := dao.NewNodeFromCreateDTO(create)

	newNode, err := s.repository.Create(ctx, newNode)
	if err != nil {
		return nodeDTO.ResponseDTO{}, err
	}

	return newNode.ToResponseDTO(), nil
}
