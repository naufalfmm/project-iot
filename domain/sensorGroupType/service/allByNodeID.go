package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (s *service) AllByNodeID(ctx echo.Context, nodeID uint64) (dao.SensorGroupTypes, error) {
	data, err := s.repository.AllByNodeID(ctx, nodeID)
	if err != nil {
		return dao.SensorGroupTypes{}, err
	}

	return data, nil
}
