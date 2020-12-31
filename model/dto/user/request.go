package user

type (
	CreateDTO struct {
		Username string
		Password string
		By       string
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
