package domain

import (
	"time"
)

type User struct {
	ID        uint64    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	IsAdmin   bool      `db:"is_admin"`
	IsActive  bool      `db:"is_active"`
	LastLogin time.Time `db:"last_login"`
	CreatedAt time.Time `db:"created_at"`
}
