package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/utils"
	"github.com/naufalfmm/project-iot/model/dao"
)

func (r *repository) Create(ctx echo.Context, newNode dao.Node) (dao.Node, error) {
	orm, err := utils.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.Node{}, err
	}

	err = orm.Create(&newNode).Error
	if err != nil {
		return dao.Node{}, err
	}

	return newNode, nil
}
