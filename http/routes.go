package http

import (
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/project-iot/common/defaultResp"
	"github.com/naufalfmm/project-iot/middleware/verifyToken"
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

	node := e.Group("/node", verifyToken.EchoMiddleware(r.Controllers.Auth.Auth))
	node.POST("", r.Controllers.Node.Create)
	node.GET("", r.Controllers.Node.All)
	node.GET("/:nodeId", r.Controllers.Node.GetByID)

	sensor := e.Group("/sensor", verifyToken.EchoMiddleware(r.Controllers.Auth.Auth))
	sensor.POST("", r.Controllers.NodeSensor.Create)
	sensor.PATCH("/:sensorId/toggle-active", r.Controllers.NodeSensor.ToggleActive)

	user := e.Group("/user")
	user.POST("/signin", r.Controllers.User.SignIn)
	user.POST("/signup", r.Controllers.User.SignUp)

	e.GET("/", func(ctx echo.Context) error {
		return defaultResp.CreateResp(ctx, r.Resource.Config.ServerName)
	})
}
