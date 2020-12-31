package dao

import (
	"time"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

type (
	User struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement"`
		Username  string    `gorm:"not null"`
		Password  string    `gorm:"not null"`
		CreatedAt time.Time `gorm:"not null"`
		CreatedBy string    `gorm:"not null"`
		UpdatedAt *time.Time
		UpdatedBy *string
	}
)

func (User) TableName() string {
	return "users"
}

func (u User) ToResponseDTO() userDTO.ResponseDTO {
	return userDTO.ResponseDTO{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
	}
}

func (u User) ToTokenResponseDTO(token string) userDTO.TokenResponseDTO {
	return userDTO.TokenResponseDTO{
		Token: token,
		User:  u.ToResponseDTO(),
	}
}

func NewUserFromCreateDTO(sur userDTO.CreateDTO) User {
	now := time.Now()

	return User{
		Username:  sur.Username,
		Password:  sur.Password,
		CreatedAt: now,
		CreatedBy: sur.By,
	}
}
