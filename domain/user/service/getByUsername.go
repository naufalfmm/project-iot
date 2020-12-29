package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	"gorm.io/gorm"
)

func (s *service) GetByUsername(ctx echo.Context, username string) (dao.User, error) {
	user, err := s.repository.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.User{}, consts.NotFoundError
		}

		return dao.User{}, err
	}

	return user, nil
}
