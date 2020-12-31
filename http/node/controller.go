package node

import (
	"github.com/naufalfmm/project-iot/handlers/node"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Controller struct {
		Node     node.Handler
		Resource resource.Resource
	}
)
