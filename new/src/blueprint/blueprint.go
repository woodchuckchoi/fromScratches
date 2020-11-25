package blueprint

import (
	"github.com/labstack/echo/v4"

	_ "github.com/woodchuckchoi/fromScratches/src/core/service"
)

func Blueprint() *echo.Echo {
	e := echo.New()

	return e
}
