package user

import (
	"errors"
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
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	err = ctx.Validate(signinReq)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	resp, err := c.User.SignIn(ctx, signinReq)
	if err != nil {
		if errors.Is(err, consts.Unauthorized) {
			return defaultResp.CreateErrorResp(ctx, http.StatusUnauthorized, consts.Unauthorized.Error())
		}

		return defaultResp.CreateErrorResp(ctx, http.StatusInternalServerError, err.Error())
	}

	return defaultResp.CreateSuccessResp(ctx, http.StatusOK, resp)
}
