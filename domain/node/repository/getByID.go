package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) GetByID(ctx echo.Context, nodeID uint64) (dao.Node, error) {
	var data dao.Node

	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.Node{}, err
	}

	err = orm.Where("id = ? AND is_deleted = ?", nodeID, false).First(&data).Error
	if err != nil {
		return dao.Node{}, err
	}

	return data, nil
}
