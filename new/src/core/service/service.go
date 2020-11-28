package service

import (
	"fmt"

	"github.com/woodchuckchoi/fromScratches/src/core/module"
)

// use functions to defind Crop's methods here or move them to where the model definition is

func init() {
	fmt.Println(module.Crop{
		URI:    "https://somewhere",
		Depth:  5,
		KeyNum: 5,
	})
}
