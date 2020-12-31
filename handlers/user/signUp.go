package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/login"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (h *handler) SignUp(ctx echo.Context, req userDTO.SignUpRequestDTO) (userDTO.TokenResponseDTO, error) {
	create := userDTO.CreateDTO{
		Username: req.Username,
		Password: req.Password,
		By:       req.Username,
	}

	newUser, err := h.domain.User.Create(ctx, create)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.TokenResponseDTO{}, err
	}

	loginData := login.ClientJWTDTO{
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	jwtToken, err := login.CreateToken(h.resource.Jwt, loginData, h.resource.Config.JwtExpiresInDuration)
	if err != nil {
		ctx.Set(consts.ResponseCode, http.StatusInternalServerError)
		return userDTO.TokenResponseDTO{}, err
	}

	return newUser.ToTokenResponseDTO(jwtToken), nil
}
