package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) GetByUsername(ctx echo.Context, username string) (dao.User, error) {
	var userData dao.User

	err := r.resource.DB.Where("username = ?", username).First(&userData).Error
	if err != nil {
		return dao.User{}, err
	}

	return userData, nil
}
