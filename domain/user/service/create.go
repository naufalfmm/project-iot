package service

import (
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/consts"
	"github.com/naufalfmm/project-iot/common/password"
	"github.com/naufalfmm/project-iot/model/dao"
	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

func (s *service) Create(ctx echo.Context, create userDTO.CreateDTO) (dao.User, error) {
	hashedPass, err := password.Hash(create.Password)
	if err != nil {
		return dao.User{}, err
	}

	create.Password = hashedPass

	newUser := dao.NewUserFromCreateDTO(create)

	newUser, err = s.repository.Create(ctx, newUser)
	if err != nil {
		switch err.(type) {
		case *pgconn.PgError:
			{
				pqErr := err.(*pgconn.PgError)
				if pqErr.Code == "23505" {
					return dao.User{}, consts.UniqueError
				}

				return dao.User{}, err
			}
		default:
			{
				return dao.User{}, err
			}
		}
	}

	newUser.Password = hashedPass

	return newUser, nil
}
