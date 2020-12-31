package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) AllByNodeID(ctx echo.Context, nodeID uint64) (dao.NodeSensors, error) {
	var data dao.NodeSensors

	err := r.resource.DB.
		Order("group_th").
		Order("created_at").
		Where("node_id = ? AND is_deleted IS FALSE", nodeID).
		Find(&data).
		Error
	if err != nil {
		return dao.NodeSensors{}, err
	}

	return data, nil
}
