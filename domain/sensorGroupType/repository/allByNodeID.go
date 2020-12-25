package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) AllByNodeID(ctx echo.Context, nodeID uint64) (dao.SensorGroupTypes, error) {
	var data dao.SensorGroupTypes

	err := r.resource.DB.
		Order("sensor_groups.th").
		Order("sensor_group_types.created_at").
		Joins("SensorGroup").
		Where("sensor_groups.node_id = ?", nodeID).
		Find(&data).
		Error
	if err != nil {
		return data, err
	}

	return data, nil
}
