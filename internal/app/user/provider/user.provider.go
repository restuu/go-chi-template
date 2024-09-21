package provider

import (
	"go-chi-template/internal/app/user"
	"go-chi-template/internal/app/user/repository"
	"go-chi-template/internal/app/user/usecase/v1"

	"github.com/google/wire"
)

var UserRepositoryProvider = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(user.UserRepository), new(*repository.UserRepository)),
)

var UserGetUsecaseProvider = wire.NewSet(
	usecase.NewUserGetUsecase,
	wire.Bind(new(user.UserGetUsecase), new(*usecase.UserGetUsecase)),
)

var UserUpsertUsecaseProvider = wire.NewSet(
	usecase.NewUserUpsertUsecase,
	wire.Bind(new(user.UserUpsertUsecase), new(*usecase.UserUpsertUsecase)),
)

var UserUsecaseProvider = wire.NewSet(
	UserRepositoryProvider,

	UserGetUsecaseProvider,
	UserUpsertUsecaseProvider,

	wire.Struct(new(UserUsecase), "*"),
	wire.Bind(new(user.UserUsecase), new(*UserUsecase)),
)

type UserUsecase struct {
	user.UserGetUsecase
	user.UserUpsertUsecase
}
