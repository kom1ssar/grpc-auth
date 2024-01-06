package model

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      int32
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
