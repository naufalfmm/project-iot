package defaultResp

import "github.com/labstack/echo/v4"

type Error struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}

func createErrorResp(ctx echo.Context, code int, err interface{}) error {
	errorData := Error{
		Code:  code,
		Error: err,
	}

	return ctx.JSON(code, errorData)
}
