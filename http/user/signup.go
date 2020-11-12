package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/naufalfmm/project-iot/common/defaultResp"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (c *Controller) SignUp(ctx echo.Context) error {
	var signupReq userDTO.SignUpRequestDTO

	err := ctx.Bind(&signupReq)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	err = ctx.Validate(signupReq)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	resp, err := c.User.SignUp(ctx, signupReq)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusInternalServerError, err.Error())
	}

	return defaultResp.CreateSuccessResp(ctx, http.StatusCreated, resp)
}
