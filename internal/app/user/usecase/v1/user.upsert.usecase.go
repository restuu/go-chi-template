package usecase

import (
	"context"

	"go-chi-template/internal/app/user"
)

type UserUpsertUsecase struct {
	userRepo user.UserRepository
}

var _ user.UserUpsertUsecase = (*UserUpsertUsecase)(nil)

func NewUserUpsertUsecase(userRepo user.UserRepository) *UserUpsertUsecase {
	return &UserUpsertUsecase{
		userRepo: userRepo,
	}
}

func (u UserUpsertUsecase) UpsertUser(ctx context.Context, user user.User) (user.User, error) {
	return u.userRepo.UpsertUser(ctx, user)
}
