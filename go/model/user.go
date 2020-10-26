package model

import (
)

type (
	User struct {
		ID		int		`json:"id"`
		Name	string	`json:"name"`
		UUID	string	`json:"uuid"`
		Low		int		`json:"low"`
		High	int		`json:"high"`
	}
)
