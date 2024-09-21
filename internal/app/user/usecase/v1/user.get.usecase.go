package usecase

import (
	"context"

	"go-chi-template/internal/app/user"
	"go-chi-template/internal/pkg/api"

	"github.com/google/uuid"
)

type UserGetUsecase struct {
	userRepo user.UserRepository
}

func NewUserGetUsecase(userRepo user.UserRepository) *UserGetUsecase {
	return &UserGetUsecase{
		userRepo: userRepo,
	}
}

func (u UserGetUsecase) GetUser(ctx context.Context, userID uuid.UUID) (user.User, error) {
	gotUser, err := u.userRepo.GetUser(ctx, userID)
	if err != nil {
		return user.User{}, api.InternalServerError(err)
	}

	if gotUser == nil {
		return user.User{}, api.NotFound()
	}

	return *gotUser, nil
}
