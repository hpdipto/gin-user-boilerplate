package user

import (
	"database/sql"
	"time"
)

// User struct type
type User struct {
	ID           uint           `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	Age          uint8          `json:"age"`
	Birthday     *time.Time     `json:"birth_day"`
	MemberNumber sql.NullString `json:"-"`
	ActivatedAt  sql.NullTime   `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
}
