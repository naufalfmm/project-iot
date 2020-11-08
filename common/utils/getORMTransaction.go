package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/resource"
	"gorm.io/gorm"
)

func GetORMTransaction(ctx echo.Context, resource resource.Resource) (*gorm.DB, error) {
	var ok bool

	orm := resource.DB

	if trxInt := ctx.Get(consts.PostgreTrx); trxInt != nil {
		if orm, ok = trxInt.(*gorm.DB); !ok {
			return nil, consts.NotTrxError
		}
	}

	return orm, nil
}
