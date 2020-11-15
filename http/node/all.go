package node

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/defaultResp"
)

func (c *Controller) All(ctx echo.Context) error {
	paramsAll := nodeDTO.NewAllRequestParamsDTO(ctx)

	all, err := c.Node.All(ctx, paramsAll)
	if err != nil {
		return defaultResp.CreateErrorResp(ctx, http.StatusInternalServerError, err.Error())
	}

	return defaultResp.CreateSuccessResp(ctx, http.StatusOK, all)
}
