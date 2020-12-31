package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/login"
)

func (h *handler) VerifyToken(ctx echo.Context, token string) (login.ClientJWTDTO, error) {
	if token == "" {
		ctx.Set(consts.ResponseCode, http.StatusUnauthorized)
		return login.ClientJWTDTO{}, consts.Unauthorized
	}

	clientDTO, err := login.DecodeToken(h.resource.Jwt, token)
	if err != nil {
		statusCode := http.StatusBadRequest
		if errors.Is(err, consts.UnclaimedToken) {
			statusCode = http.StatusUnauthorized
		}
		ctx.Set(consts.ResponseCode, statusCode)
		return login.ClientJWTDTO{}, err
	}

	_, err = h.domain.User.GetByUsername(ctx, clientDTO.Username)
	if err != nil {
		if errors.Is(err, consts.NotFoundError) {
			ctx.Set(consts.ResponseCode, http.StatusUnauthorized)
			return login.ClientJWTDTO{}, consts.Unauthorized
		}

		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return login.ClientJWTDTO{}, err
	}

	return clientDTO, nil
}
