package dao

import (
	"time"

	userDTO "github.com/naufalfmm/project-iot/model/dto/user"
)

type (
	User struct {
		ID        uint64     `gorm:"PRIMARY_KEY`
		Username  string     `gorm:"not null"`
		Password  string     `gorm:"not null"`
		CreatedAt time.Time  `gorm:"not null"`
		UpdatedAt *time.Time `gorm:"not null"`
	}
)

func (User) TableName() string {
	return "users"
}

func (u User) ToResponseDTO() userDTO.ResponseDTO {
	return userDTO.ResponseDTO(u)
}

func (u User) ToSignUpResponseDTO() userDTO.SignUpResponseDTO {
	return userDTO.SignUpResponseDTO{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUserFromSignUpRequestDTO(sur userDTO.SignUpRequestDTO) User {
	now := time.Now()

	return User{
		Username:  sur.Username,
		Password:  sur.Password,
		CreatedAt: now,
	}
}
