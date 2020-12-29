package node

import (
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"

	"github.com/naufalfmm/project-iot/common/defaultResp"
	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (c *Controller) Create(ctx echo.Context) error {
	var bodyReq nodeDTO.CreateRequestBodyDTO

	if err := ctx.Bind(&bodyReq); err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}
	if err := ctx.Validate(&bodyReq); err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	loginData, err := c.getCurrentLogin(ctx)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	createReq := nodeDTO.CreateRequestDTO{
		Body: bodyReq,
		By:   loginData,
	}

	resp, err := c.Node.Create(ctx, createReq)
	if err != nil {
		return defaultResp.CreateResp(ctx, err.Error())
	}

	ctx.Set(consts.ResponseCode, http.StatusCreated)
	return defaultResp.CreateResp(ctx, resp)
}
