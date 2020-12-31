package user

import "time"

type (
	ResponseDTO struct {
		ID        uint64    `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
		CreatedBy string    `json:"created_by"`
	}
)

type (
	TokenResponseDTO struct {
		Token string      `json:"token"`
		User  ResponseDTO `json:"user"`
	}
)
