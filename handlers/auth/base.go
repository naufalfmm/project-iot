package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/login"
	"github.com/naufalfmm/project-iot/domain"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Handler interface {
		VerifyToken(ctx echo.Context, token string) (login.ClientJWTDTO, error)
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
