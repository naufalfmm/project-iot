package nodeSensor

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/defaultResp"
	"github.com/naufalfmm/project-iot/common/login"
	nodeSensorDTO "github.com/naufalfmm/project-iot/model/dto/nodeSensor"
)

func (c *Controller) Create(ctx echo.Context) error {
	var bodyReq nodeSensorDTO.CreateRequestBodyDTO

	if err := ctx.Bind(&bodyReq); err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	if err := ctx.Validate(&bodyReq); err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	loginData, err := login.GetCurrentLogin(ctx)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusBadRequest)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	createReq := nodeSensorDTO.CreateRequestDTO{
		Body: bodyReq,
		By:   loginData,
	}

	resp, err := c.NodeSensor.Create(ctx, createReq)
	if err != nil {
		return defaultResp.CreateResp(ctx, err.Error())
	}

	ctx.Set(consts.ResponseCode, http.StatusCreated)
	return defaultResp.CreateResp(ctx, resp)
}
