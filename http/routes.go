package http

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/login"
	"github.com/naufalfmm/project-iot/resource"
)

type (
	Routes struct {
		Controllers Controllers
		Resource    resource.Resource
	}
)

func (r *Routes) Register(e *echo.Echo) {
	sensorData := e.Group("/sensorData")
	sensorData.GET("/post", r.Controllers.SensorData.CreateFromNode)

	node := e.Group("/node", login.EchoMiddleware(r.Resource.Jwt))
	node.POST("", r.Controllers.Node.Create)
}
