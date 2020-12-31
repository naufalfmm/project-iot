package user

import (
	"errors"
	"net/http"

	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/password"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/login"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (h *handler) SignIn(ctx echo.Context, req userDTO.SignInRequestDTO) (userDTO.TokenResponseDTO, error) {
	userData, err := h.domain.User.GetByUsername(ctx, req.Username)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.NotFoundError) {
			statusCode = http.StatusUnauthorized
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return userDTO.TokenResponseDTO{}, err
	}

	valid, err := password.Verify(userData.Password, req.Password)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.TokenResponseDTO{}, err
	}

	if !valid {
		ctx.Set(consts.ResponseCode, http.StatusUnauthorized)
		return userDTO.TokenResponseDTO{}, consts.Unauthorized
	}

	loginData := login.ClientJWTDTO{
		ID:       userData.ID,
		Username: userData.Username,
	}

	jwtToken, err := login.CreateToken(h.resource.Jwt, loginData, h.resource.Config.JwtExpiresInDuration)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.TokenResponseDTO{}, err
	}

	return userData.ToTokenResponseDTO(jwtToken), nil
}
