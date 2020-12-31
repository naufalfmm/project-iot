package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	"gorm.io/gorm"
)

func (s *service) GetByID(ctx echo.Context, nodeID uint64) (dao.Node, error) {
	node, err := s.repository.GetByID(ctx, nodeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Node{}, consts.NotFoundError
		}

		return dao.Node{}, err
	}

	return node, nil
}
