package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
)

func (s *service) ToggleActive(ctx echo.Context, sensorID uint64) error {
	rowsAff, err := s.repository.ToggleActive(ctx, sensorID)
	if err != nil {
		return err
	}

	if rowsAff < 1 {
		return consts.NotFoundError
	}

	return nil
}
