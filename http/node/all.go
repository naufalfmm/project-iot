package node

import (
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/defaultResp"

	nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"
)

func (c *Controller) All(ctx echo.Context) error {
	paramsAll := nodeDTO.NewAllRequestParamsDTO(ctx)

	all, err := c.Node.All(ctx, paramsAll)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return defaultResp.CreateResp(ctx, err.Error())
	}

	return defaultResp.CreateResp(ctx, all)
}
