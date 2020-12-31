package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (s *service) GetUniqueByNodeID(ctx echo.Context, nodeID uint64) (dao.GroupSensors, error) {
	return s.repository.GetUniqueByNodeID(ctx, nodeID)
}
