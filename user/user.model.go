package user

import (
	"database/sql"
	"time"
)

// User struct type
type User struct {
	ID           uint           `json:"id"`
	FirstName    string         `json:"first_name" binding:"required"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email" binding:"required"`
	Password     string         `json:"password" binding:"required"`
	Age          uint8          `json:"age"`
	Birthday     *time.Time     `json:"birth_day" `
	MemberNumber sql.NullString `json:"-"`
	ActivatedAt  sql.NullTime   `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
}

// Info struct type
type Info struct {
	ID           uint           `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	Age          uint8          `json:"age"`
	Birthday     *time.Time     `json:"birth_day" `
	MemberNumber sql.NullString `json:"-"`
	ActivatedAt  sql.NullTime   `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
}

// LoginUser struct type
type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
