package user

import (
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/defaultResp"

	"github.com/labstack/echo/v4"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (c *Controller) SignIn(ctx echo.Context) error {
	var signinReq userDTO.SignInRequestDTO

	err := ctx.Bind(&signinReq)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	err = ctx.Validate(signinReq)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	resp, err := c.User.SignIn(ctx, signinReq)
	if err != nil {
		return defaultResp.CreateResp(ctx, err.Error())
	}

	return defaultResp.CreateResp(ctx, resp)
}
