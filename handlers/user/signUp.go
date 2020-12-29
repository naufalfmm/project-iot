package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/login"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (h *handler) SignUp(ctx echo.Context, req userDTO.SignUpRequestDTO) (userDTO.SignUpTokenResponseDTO, error) {
	newUser, err := h.domain.User.SignUp(ctx, req)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.SignUpTokenResponseDTO{}, err
	}

	loginData := login.ClientJWTDTO{
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	jwtToken, err := login.CreateToken(h.resource.Jwt, loginData, h.resource.Config.JwtExpiresInDuration)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.SignUpTokenResponseDTO{}, err
	}

	return newUser.ToSignUpTokenResponseDTO(jwtToken), nil
}
