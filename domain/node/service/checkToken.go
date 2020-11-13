package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"gorm.io/gorm"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (s *service) CheckToken(ctx echo.Context, token string) (nodeDTO.ResponseDTO, error) {
	data, err := s.repository.GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nodeDTO.ResponseDTO{}, consts.Unauthorized
		}

		return nodeDTO.ResponseDTO{}, err
	}

	return data.ToResponseDTO(), nil
}
