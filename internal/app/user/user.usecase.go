package user

import (
	"context"

	"github.com/google/uuid"
)

type UserUpsertUsecase interface {
	UpsertUser(ctx context.Context, user User) (User, error)
}

type UserGetUsecase interface {
	GetUser(ctx context.Context, userID uuid.UUID) (User, error)
}

type UserUsecase interface {
	UserUpsertUsecase
	UserGetUsecase
}
