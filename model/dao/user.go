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
