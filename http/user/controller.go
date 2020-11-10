package user

import (
	"github.com/naufalfmm/project-iot/handlers/user"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		User     user.Handler
		Resource resource.Resource
	}
)
