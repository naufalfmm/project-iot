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
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = ctx.Validate(signinReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	resp, err := c.User.SignIn(ctx, signinReq)
	if err != nil {
		if errors.Is(err, consts.Unauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized, consts.Unauthorized)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, defaultResp.CreateSuccessResp(http.StatusOK, resp))
}
