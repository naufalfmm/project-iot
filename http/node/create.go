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
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(&bodyReq); err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	loginData, err := c.getCurrentLogin(ctx)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusBadRequest, err.Error())
	}

	createReq := nodeDTO.CreateRequestDTO{
		Body: bodyReq,
		By:   loginData,
	}

	resp, err := c.Node.Create(ctx, createReq)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusInternalServerError, err.Error())
	}

	return defaultResp.CreateSuccessResp(ctx, http.StatusCreated, resp)
}
