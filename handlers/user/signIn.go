package user

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/login"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (h *handler) SignIn(ctx echo.Context, req userDTO.SignInRequestDTO) (userDTO.SignInTokenResponseDTO, error) {
	userData, err := h.domain.User.SignIn(ctx, req)
	if err != nil {
		return userDTO.SignInTokenResponseDTO{}, err
	}

	loginData := login.ClientJWTDTO{
		ID:       userData.ID,
		Username: userData.Username,
	}

	jwtToken, err := login.CreateToken(h.resource.Jwt, loginData, h.resource.Config.JwtExpiresInDuration)
	if err != nil {
		return userDTO.SignInTokenResponseDTO{}, err
	}

	return userData.ToSignInTokenResponseDTO(jwtToken), nil
}
