package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	Info      *UserInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
	Role        string `db:"role"`
}
