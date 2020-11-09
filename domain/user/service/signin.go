package service

import (
	"errors"

	"github.com/naufalfmm/project-iot/common/password"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"gorm.io/gorm"

	"github.com/naufalfmm/project-iot/model/dao"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (s *service) SignIn(ctx echo.Context, signInData userDTO.SignInRequestDTO) (userDTO.ResponseDTO, error) {
	whereUser := dao.User{
		Username: signInData.Username,
	}

	userData, err := s.repository.Get(ctx, whereUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDTO.ResponseDTO{}, consts.Unauthorized
		}

		return userDTO.ResponseDTO{}, err
	}

	valid, err := password.Verify(userData.Password, signInData.Password)
	if err != nil {
		return userDTO.ResponseDTO{}, err
	}

	if !valid {
		return userDTO.ResponseDTO{}, consts.Unauthorized
	}

	return userData.ToResponseDTO(), nil
}
