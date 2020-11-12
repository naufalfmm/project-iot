package resource

import (
	"github.com/naufalfmm/project-iot/resource/config"
	"github.com/naufalfmm/project-iot/resource/jwt"
	"github.com/naufalfmm/project-iot/resource/validator"
	"gorm.io/gorm"
)

type Resource struct {
	Config    *config.EnvConfig
	DB        *gorm.DB
	Jwt       jwt.JWT
	Validator *validator.CustomValidator
}

func New(config *config.EnvConfig, db *gorm.DB, jwt jwt.JWT, validator *validator.CustomValidator) Resource {
	return Resource{
		Config:    config,
		DB:        db,
		Jwt:       jwt,
		Validator: validator,
	}
}
