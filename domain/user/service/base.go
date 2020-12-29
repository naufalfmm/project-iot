package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/user/repository"
	"github.com/naufalfmm/project-iot/model/dao"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		SignUp(ctx echo.Context, signUpData userDTO.SignUpRequestDTO) (dao.User, error)
		GetByUsername(ctx echo.Context, username string) (dao.User, error)
	}

	service struct {
		resource   resource.Resource
		repository repository.Repository
	}
)

func New(resource resource.Resource, repository repository.Repository) (Service, error) {
	return &service{
		resource:   resource,
		repository: repository,
	}, nil
}
