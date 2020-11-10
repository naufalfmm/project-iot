package login

import (
	"net/http"
	"strings"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/resource/jwt"
)

type ClientClaims struct {
	*jwtGo.StandardClaims
	Data ClientJWTDTO `json:"data"`
}

type ClientJWTDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func DecodeToken(j jwt.JWT, token string) (ClientJWTDTO, error) {
	var clientClaims ClientClaims

	claims, err := j.Parse(token, &clientClaims)
	if err != nil {
		return ClientJWTDTO{}, err
	}

	clientClaims, ok := claims.(ClientClaims)
	if !ok {
		return ClientJWTDTO{}, consts.WrongClaimsFormat
	}

	return clientClaims.Data, nil
}

func CreateToken(j jwt.JWT, data ClientJWTDTO, exp time.Duration) (string, error) {
	issuedAt := time.Now()

	clientClaims := ClientClaims{
		&jwtGo.StandardClaims{
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
				return echo.NewHTTPError(http.StatusUnauthorized, consts.Unauthorized.Error())
			}

			clientDTO, err := DecodeToken(j, tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			ctx.Set(consts.UserLoginKey, clientDTO)
			return next(ctx)
		}
	}
}
