package login

import (
	"time"

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

func GetCurrentLogin(ctx echo.Context) (ClientJWTDTO, error) {
	loginDTO, ok := ctx.Get(consts.UserLoginKey).(ClientJWTDTO)
	if !ok {
		err := consts.UndefinedLoginDataError
		return ClientJWTDTO{}, err
	}

	return loginDTO, nil
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
