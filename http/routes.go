package http

import (
	"net/http"

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
	sensorData := e.Group("/sensorData")
	sensorData.GET("/post", r.Controllers.SensorData.CreateFromNode)

	node := e.Group("/node", login.EchoMiddleware(r.Resource.Jwt))
	node.POST("", r.Controllers.Node.Create)

	user := e.Group("/user")
	user.POST("/login", r.Controllers.User.SignIn)
	user.POST("/signup", r.Controllers.User.SignUp, login.EchoMiddleware(r.Resource.Jwt))

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, defaultResp.CreateSuccessResp(http.StatusOK, r.Resource.Config.ServerName))
	})
}
