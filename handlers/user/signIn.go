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

func (h *handler) SignIn(ctx echo.Context, req userDTO.SignInRequestDTO) (userDTO.SignInTokenResponseDTO, error) {
	userData, err := h.domain.User.GetByUsername(ctx, req.Username)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, consts.NotFoundError) {
			statusCode = http.StatusUnauthorized
		}

		ctx.Set(consts.ResponseCode, statusCode)
		return userDTO.SignInTokenResponseDTO{}, err
	}

	valid, err := password.Verify(userData.Password, signInData.Password)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.SignInTokenResponseDTO{}, err
	}

	if !valid {
		ctx.Set(consts.ResponseCode, http.StatusUnauthorized)
		return userDTO.SignInTokenResponseDTO{}, consts.Unauthorized
	}

	loginData := login.ClientJWTDTO{
		ID:       userData.ID,
		Username: userData.Username,
	}

	jwtToken, err := login.CreateToken(h.resource.Jwt, loginData, h.resource.Config.JwtExpiresInDuration)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.SignInTokenResponseDTO{}, err
	}

	return userData.ToSignInTokenResponseDTO(jwtToken), nil
}
