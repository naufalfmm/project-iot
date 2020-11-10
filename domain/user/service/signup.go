package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/model/dao"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (s *service) SignUp(ctx echo.Context, signUpData userDTO.SignUpRequestDTO) (userDTO.SignUpResponseDTO, error) {
	newUser := dao.NewUserFromSignUpRequestDTO(signUpData)

	newUser, err := s.repository.Create(ctx, newUser)
	if err != nil {
		return userDTO.SignUpResponseDTO{}, err
	}

	return newUser.ToSignUpResponseDTO(), nil
}