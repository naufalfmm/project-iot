package defaultResp

import "github.com/labstack/echo/v4"

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func CreateSuccessResp(ctx echo.Context, code int, data interface{}) error {
	successData := Success{
		Code: code,
		Data: data,
	}

	return ctx.JSON(code, successData)
}
