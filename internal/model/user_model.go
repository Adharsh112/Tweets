package model

import "time"

type (
	userModel struct {
		ID        int64
		Email     string
		username  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
