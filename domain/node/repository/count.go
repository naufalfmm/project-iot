package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Count(ctx echo.Context) (int64, error) {
	var count int64

	err := r.resource.DB.Model(&dao.Node{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
