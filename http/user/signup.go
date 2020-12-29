package user

import (
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"

	"github.com/naufalfmm/project-iot/common/defaultResp"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (c *Controller) SignUp(ctx echo.Context) error {
	var signupReq userDTO.SignUpRequestDTO

	err := ctx.Bind(&signupReq)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	err = ctx.Validate(signupReq)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	resp, err := c.User.SignUp(ctx, signupReq)
	if err != nil {
		return defaultResp.CreateResp(ctx, err.Error())
	}

	ctx.Set(consts.ResponseCode, http.StatusCreated)
	return defaultResp.CreateResp(ctx, resp)
}
