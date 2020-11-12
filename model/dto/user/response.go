package user

import "time"

type (
	ResponseDTO struct {
		ID        uint64     `json:"id"`
		Username  string     `json:"username"`
		Password  string     `json:"password"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}
)

type (
	SignUpResponseDTO struct {
		ID        uint64     `json:"id"`
		Username  string     `json:"username"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	SignUpTokenResponseDTO struct {
		Token string            `json:"token"`
		User  SignUpResponseDTO `json:"user"`
	}

	SignInResponseDTO struct {
		ID        uint64     `json:"id"`
		Username  string     `json:"username"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	SignInTokenResponseDTO struct {
		Token string            `json:"token"`
		User  SignInResponseDTO `json:"user"`
	}
)

func (su SignUpResponseDTO) ToSignUpTokenResponseDTO(token string) SignUpTokenResponseDTO {
	return SignUpTokenResponseDTO{
		Token: token,
		User:  su,
	}
}

func (si SignInResponseDTO) ToSignInTokenResponseDTO(token string) SignInTokenResponseDTO {
	return SignInTokenResponseDTO{
		Token: token,
		User:  si,
	}
}
