package node

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/defaultResp"
)

func (c *Controller) GetByID(ctx echo.Context) error {
	nodeID, err := strconv.ParseUint(ctx.Param("nodeId"), 10, 64)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return defaultResp.CreateResp(ctx, err)
	}

	data, err := c.Node.GetByID(ctx, nodeID)
	if err != nil {
		return defaultResp.CreateResp(ctx, err)
	}

	return defaultResp.CreateResp(ctx, data)
}
