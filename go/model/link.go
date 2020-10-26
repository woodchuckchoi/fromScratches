package model

import (
)

type (
	struct Link {
		ID		int		`json:"id"`
		UserID	int		`json:"user_id"`
		Link	string	`json:"link"`
	}
)
