package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	FullName string    `json:"full_name" db:"full_name"`
}
