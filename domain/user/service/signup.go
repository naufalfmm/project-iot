package service

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/password"
	"github.com/naufalfmm/project-iot/model/dao"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (s *service) SignUp(ctx echo.Context, signUpData userDTO.SignUpRequestDTO) (dao.User, error) {
	hashedPass, err := password.Hash(signUpData.Password)
	if err != nil {
		return dao.User{}, err
	}

	signUpData.Password = hashedPass

	newUser := dao.NewUserFromSignUpRequestDTO(signUpData)

	newUser, err = s.repository.Create(ctx, newUser)
	if err != nil {
		return dao.User{}, err
	}

	newUser.Password = hashedPass

	return newUser, nil
}
