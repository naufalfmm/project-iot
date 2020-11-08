package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/model/dao"
	"gorm.io/gorm"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) CheckToken(ctx echo.Context, token string) (nodeDTO.ResponseDTO, error) {
	whereQuery := dao.Node{
		Token: token,
	}

	data, err := s.repository.Get(ctx, whereQuery)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nodeDTO.ResponseDTO{}, consts.Unauthorized
		}

		return nodeDTO.ResponseDTO{}, err
	}

	return data.ToResponseDTO(), nil
}
