package resource

import (
	"github.com/naufalfmm/project-iot/resource/config"
	"github.com/naufalfmm/project-iot/resource/jwt"
	"gorm.io/gorm"
)

type Resource struct {
	Config *config.EnvConfig
	DB     *gorm.DB
	Jwt    jwt.JWT
}

func New(config *config.EnvConfig, db *gorm.DB, jwt jwt.JWT) Resource {
	return Resource{
		Config: config,
		DB:     db,
		Jwt:    jwt,
	}
}
