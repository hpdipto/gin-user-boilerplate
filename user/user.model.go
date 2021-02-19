package user

import (
	"database/sql"
	"time"
)

// User struct type
type User struct {
	ID           uint
	FirstName    string
	LastName     string
	Email        string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
