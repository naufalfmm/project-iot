package login

import (
	"net/http"
	"strings"
	"time"

	"github.com/naufalfmm/project-iot/common/defaultResp"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/resource/jwt"
)

type ClientClaims struct {
	jwtGo.StandardClaims
	Data ClientJWTDTO `json:"data"`
}

type ClientJWTDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func DecodeToken(j jwt.JWT, token string) (ClientJWTDTO, error) {
	var clientClaims ClientClaims

	_, err := j.Parse(token, &clientClaims)
	if err != nil {
		return ClientJWTDTO{}, err
	}

	return clientClaims.Data, nil
}

func CreateToken(j jwt.JWT, data ClientJWTDTO, exp time.Duration) (string, error) {
	issuedAt := time.Now()

	clientClaims := ClientClaims{
		jwtGo.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: issuedAt.Add(exp).Unix(),
		},
		data,
	}

	return j.Create(&clientClaims)
}

func EchoMiddleware(j jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			tokenString := ctx.Request().Header.Get("Authorization")

			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			if tokenString == "" {
				ctx.Set(consts.ResponseCode, http.StatusUnauthorized)
				return defaultResp.CreateResp(ctx, consts.Unauthorized.Error())
			}

			clientDTO, err := DecodeToken(j, tokenString)
			if err != nil {
				ctx.Set(consts.ResponseCode, http.StatusBadRequest)
				return defaultResp.CreateResp(ctx, err.Error())
			}

			ctx.Set(consts.UserLoginKey, clientDTO)
			return next(ctx)
		}
	}
}
