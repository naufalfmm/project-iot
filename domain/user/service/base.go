package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/domain/user/repository"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Service interface {
		SignIn(ctx echo.Context, signInData userDTO.SignInRequestDTO) (userDTO.SignInResponseDTO, error)
		SignUp(ctx echo.Context, signUpData userDTO.SignUpRequestDTO) (userDTO.SignUpResponseDTO, error)
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
