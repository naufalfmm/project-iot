package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Get(ctx echo.Context, whereQuery dao.User) (dao.User, error) {
	var userData dao.User

	err := r.resource.DB.First(&userData, whereQuery).Error
	if err != nil {
		return dao.User{}, nil
	}

	return userData, err
}
