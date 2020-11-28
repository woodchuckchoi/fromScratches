package blueprint

import (
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/woodchuckchoi/fromScratches/docs"
	_ "github.com/woodchuckchoi/fromScratches/src/core/service"
)

var (
	E                echo.Echo
	ControllerGroup1 *echo.Group
)

func init() {
	// Swagger initiation
	E.GET("/docs/*", echoSwagger.WrapHandler)

	ControllerGroup1 = E.Group("/api/v1")
}
