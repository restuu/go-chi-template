// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"go-chi-template/internal/app/user/provider"
	"go-chi-template/internal/app/user/repository"
	"go-chi-template/internal/app/user/usecase/v1"
)

// Injectors from wire.go:

func initApp() (*app, error) {
	sql, err := connectDB()
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(sql)
	userGetUsecase := usecase.NewUserGetUsecase(userRepository)
	userUpsertUsecase := usecase.NewUserUpsertUsecase(userRepository)
	userUsecase := &provider.UserUsecase{
		UserGetUsecase:    userGetUsecase,
		UserUpsertUsecase: userUpsertUsecase,
	}
	mainUsecases := usecases{
		UserUsecase: userUsecase,
	}
	router := initRouter(mainUsecases)
	server := initServer(router)
	mainApp := &app{
		srv:      server,
		usecases: mainUsecases,
	}
	return mainApp, nil
}
