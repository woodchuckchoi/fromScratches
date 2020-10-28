package route

import (
	"github.com/labstack/echo/v4"
	"github.com/woodchuckchoi/sweetpet/handler"
)

func Match(e *echo.Echo, h *handler.Handler) {
	// e.GET("/api/v1/health/:link", h.RetrieveAllHealthEntries)
	e.GET("/api/v1/health/:link", h.RetrieveRangedHealthEntries)
	e.POST("/api/v1/user/create", h.Register)
	e.PUT("/api/v1/user/modify", h.ModifyThreshold)
	e.GET("/api/v1/user/new_link", h.GenerateLink)

}
