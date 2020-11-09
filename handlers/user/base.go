package user

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Handler interface {
		SignIn(ctx echo.Context, req userDTO.SignInRequestDTO) (userDTO.ResponseDTO, error)
		SignUp(ctx echo.Context, req userDTO.SignUpRequestDTO) (userDTO.SignUpTokenResponseDTO, error)
	}
	handler struct {
		domain   domain.Domain
		resource resource.Resource
	}
)

func NewHandler(domain domain.Domain, resource resource.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
