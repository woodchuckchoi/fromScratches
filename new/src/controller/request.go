package controller

import (
	"github.com/labstack/echo/v4"
)

func requestAnalytics(c echo.Context) error {

	uris := c.FormValue("uris")
	keys := c.FormValue("keys")
	deps := c.FormValue("deps")

	if func(args ...[]interface{}) error {
		if len(args) == 0 {
			return error
		}

		cnt := len(args[0])

		for idx := 1; idx < len(args); idx++ {
			if cnt != len(args[idx]) {
				return error
			}
		}

		return nil
	}(uris, keys, deps) != nil {

	}

}
