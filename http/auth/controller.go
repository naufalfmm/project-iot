package auth

import (
	"github.com/naufalfmm/project-iot/handlers/auth"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		Auth     auth.Handler
		Resource resource.Resource
	}
)
