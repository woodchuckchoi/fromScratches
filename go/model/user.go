package model

import ()

type (
	User struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name"`
		UUID string `json:"uuid,omitempty"`
		Low  int    `json:"low,omitempty"`
		High int    `json:"high,omitempty"`
		Link string `json:"link,omitempty"`
	}
)
