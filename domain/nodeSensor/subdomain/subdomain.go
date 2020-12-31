package subdomain

import (
	groupSensor "github.com/naufalfmm/project-iot/domain/nodeSensor/subdomain/groupSensor/service"
)

type (
	Subdomain struct {
		GroupSensor groupSensor.Service
	}
)
