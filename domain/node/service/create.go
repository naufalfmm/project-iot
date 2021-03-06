package service

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) Create(ctx echo.Context, create nodeDTO.CreateDTO) (dao.Node, error) {
	newNodeDTO := dao.NewNodeFromCreateDTO(create)

	newNode, err := s.repository.Create(ctx, newNodeDTO)
	if err != nil {
		switch err.(type) {
		case *pgconn.PgError:
			{
				pqErr := err.(*pgconn.PgError)
				if pqErr.Code == "23505" {
					return dao.Node{}, consts.UniqueError
				}

				return dao.Node{}, err
			}
		default:
			{
				return dao.Node{}, err
			}
		}
	}

	return newNode, nil
}
