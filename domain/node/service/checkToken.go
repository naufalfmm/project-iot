package service

import (
	"errors"

	"github.com/naufalfmm/project-iot/model/dao"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"gorm.io/gorm"
)

func (s *service) CheckToken(ctx echo.Context, token string) (dao.Node, error) {
	data, err := s.repository.GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.Node{}, consts.Unauthorized
		}

		return dao.Node{}, err
	}

	return data, nil
}
