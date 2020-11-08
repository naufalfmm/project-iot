package node

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/naufalfmm/project-iot/common/defaultResp"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (c *Controller) Create(ctx echo.Context) error {
	var bodyReq nodeDTO.CreateRequestBodyDTO

	if err := ctx.Bind(&bodyReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(&bodyReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	loginData, err := c.getCurrentLogin(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	createReq := nodeDTO.CreateRequestDTO{
		Body: bodyReq,
		By:   loginData,
	}

	resp, err := c.Node.Create(ctx, createReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, defaultResp.CreateSuccessResp(http.StatusCreated, resp))
}
