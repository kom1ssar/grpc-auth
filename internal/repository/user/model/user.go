package model

import "time"

type User struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      int32     `db:"role"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
