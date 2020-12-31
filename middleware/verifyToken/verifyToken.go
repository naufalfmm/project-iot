package verifyToken

import (
	"strings"

	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/defaultResp"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/handlers/auth"
)

func EchoMiddleware(h auth.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			tokenString := ctx.Request().Header.Get("Authorization")

			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			clientData, err := h.VerifyToken(ctx, tokenString)
			if err != nil {
				return defaultResp.CreateResp(ctx, err.Error())
			}

			ctx.Set(consts.UserLoginKey, clientData)
			return next(ctx)
		}
	}
}
