package controller

import (
	"github.com/woodchuckchoi/fromScratches/src/blueprint"
)

func init() {
	blueprint.ControllerGroup1.POST("/request", requestAnalytics) // requestAnalytics is private to this package, since its first letter is under-case
}
