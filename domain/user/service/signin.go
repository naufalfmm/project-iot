package service

import (
	"errors"

	"github.com/naufalfmm/project-iot/common/password"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"gorm.io/gorm"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (s *service) SignIn(ctx echo.Context, signInData userDTO.SignInRequestDTO) (userDTO.SignInResponseDTO, error) {
	userData, err := s.repository.GetByUsername(ctx, signInData.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDTO.SignInResponseDTO{}, consts.Unauthorized
		}

		return userDTO.SignInResponseDTO{}, err
	}

	valid, err := password.Verify(userData.Password, signInData.Password)
	if err != nil {
		return userDTO.SignInResponseDTO{}, err
	}

	if !valid {
		return userDTO.SignInResponseDTO{}, consts.Unauthorized
	}

	return userData.ToSignInResponseDTO(), nil
}
