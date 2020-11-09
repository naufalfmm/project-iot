package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Create(ctx echo.Context, newUser dao.User) (dao.User, error) {
	err := r.resource.DB.Create(&newUser).Error
	if err != nil {
		return dao.User{}, err
	}

	return newUser, nil
}
