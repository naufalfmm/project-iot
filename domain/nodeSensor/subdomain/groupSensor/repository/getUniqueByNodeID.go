package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) GetUniqueByNodeID(ctx echo.Context, nodeID uint64) (dao.GroupSensors, error) {
	var groups dao.GroupSensors

	err := r.resource.DB.
		Where("node_id = ? AND is_deleted = ?", nodeID, false).
		Select("DISTINCT ON(group_th) group_th, group_label, node_id").
		Find(&groups).
		Error
	if err != nil {
		return dao.GroupSensors{}, err
	}

	return groups, nil
}
