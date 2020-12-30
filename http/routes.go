package http

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/defaultResp"
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
	sensorData := e.Group("/sensor-data")
	sensorData.GET("/post", r.Controllers.SensorData.CreateFromNode)

	node := e.Group("/node", login.EchoMiddleware(r.Resource.Jwt))
	node.POST("", r.Controllers.Node.Create)
	node.GET("", r.Controllers.Node.All)

	user := e.Group("/user")
	user.POST("/signin", r.Controllers.User.SignIn)
	user.POST("/signup", r.Controllers.User.SignUp)

	e.GET("/", func(ctx echo.Context) error {
		return defaultResp.CreateResp(ctx, r.Resource.Config.ServerName)
	})
}
