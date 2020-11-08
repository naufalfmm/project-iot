package node

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/login"

	"github.com/naufalfmm/project-iot/handlers/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		Node     node.Handler
		Resource resource.Resource
	}
)

func (c *Controller) getCurrentLogin(ctx echo.Context) (login.ClientJWTDTO, error) {
	loginDTO, ok := ctx.Get(consts.UserLoginKey).(login.ClientJWTDTO)
	if !ok {
		err := consts.UndefinedLoginDataError
		return login.ClientJWTDTO{}, err
	}

	return loginDTO, nil
}
