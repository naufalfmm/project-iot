package user

import "time"

type (
	CreateRequestDTO struct {
		ID        uint64     `json:"id"`
		Username  string     `json:"username"`
		Password  string     `json:"password"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}
)

type (
	SignUpRequestDTO struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignInRequestDTO struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
