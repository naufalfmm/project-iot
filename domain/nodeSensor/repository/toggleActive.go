package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
	"gorm.io/gorm"
)

func (r *repository) ToggleActive(ctx echo.Context, sensorID uint64) (int64, error) {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return 0, err
	}

	orm = orm.Model(&dao.NodeSensor{}).Where("id = ?", sensorID).Update("is_active", gorm.Expr("NOT is_active"))
	if err != nil {
		return 0, err
	}

	return orm.RowsAffected, nil
}
