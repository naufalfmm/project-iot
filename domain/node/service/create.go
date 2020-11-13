package service

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) Create(ctx echo.Context, create nodeDTO.CreateDTO) (nodeDTO.ResponseDTO, error) {
	newNode := dao.NewNodeFromCreateDTO(create)

	newNode, err := s.repository.Create(ctx, newNode)
	if err != nil {
		switch err.(type) {
		case *pgconn.PgError:
			{
				pqErr := err.(*pgconn.PgError)
				if pqErr.Code == "23505" {
					return nodeDTO.ResponseDTO{}, consts.UniqueError
				} else {
					return nodeDTO.ResponseDTO{}, err
				}
			}
		default:
			{
				return nodeDTO.ResponseDTO{}, err
			}
		}
	}

	return newNode.ToResponseDTO(), nil
}
