package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt time.Time
}
