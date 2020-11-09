package user

import (
	"github.com/labstack/echo/v4"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (h *handler) SignIn(ctx echo.Context, req userDTO.SignInRequestDTO) (userDTO.ResponseDTO, error) {
	return h.domain.User.SignIn(ctx, req)
}
