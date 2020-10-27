package model

import ()

type (
	User struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
		Low  int    `json:"low,omitempty"`
		High int    `json:"high,omitempty"`
	}
)
