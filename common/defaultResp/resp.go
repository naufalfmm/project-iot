package defaultResp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
)

func CreateResp(ctx echo.Context, data interface{}) error {
	respCodeCtx := ctx.Get(consts.ResponseCode)

	if respCodeCtx == nil {
		return createSuccessResp(ctx, http.StatusOK, data)
	}

	respCode := respCodeCtx.(int)
	if respCode < 300 {
		return createSuccessResp(ctx, respCode, data)
	}

	return createErrorResp(ctx, respCode, data)
}
