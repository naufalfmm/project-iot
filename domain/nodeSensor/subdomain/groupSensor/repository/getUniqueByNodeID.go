package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) GetUniqueByNodeID(ctx echo.Context, nodeID uint64) (dao.GroupSensors, error) {
	var groups dao.GroupSensors

	err := r.resource.DB.
		Where("node_id = ?", nodeID).
		Distinct("group_th").
		Find(&groups).
		Error
	if err != nil {
		return dao.GroupSensors{}, err
	}

	return groups, nil
}
