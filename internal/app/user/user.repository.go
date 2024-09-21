package user

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	UpsertUser(ctx context.Context, user User) (User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
}
