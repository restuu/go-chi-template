package http

import (
	"encoding/json"
	"net/http"

	"go-chi-template/internal/app/user"
	"go-chi-template/internal/pkg/api"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type userHTTPHandler struct {
	userUsecase user.UserUsecase
}

func SetUserHTTPHandler(r chi.Router, userUsecase user.UserUsecase) {
	h := &userHTTPHandler{
		userUsecase: userUsecase,
	}

	r.Route("/v1/users", func(r chi.Router) {
		// Without auth
		r.Group(func(r chi.Router) {
			r.Post("/register", api.Handle(h.register))
		})

		// With auth
		r.Group(func(r chi.Router) {
			// TODO: decode auth middleware
			r.Use()

			r.Get("/me", api.Handle(h.getCurrentUser))
		})
	})
}

func (h userHTTPHandler) register(w http.ResponseWriter, r *http.Request) error {
	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return api.InternalServerError(err)
	}

	ctx := r.Context()

	gotUser, err := h.userUsecase.UpsertUser(ctx, u)
	if err != nil {
		return err
	}

	api.WriteSuccess(ctx, w, gotUser)

	return nil
}

func (h userHTTPHandler) getCurrentUser(w http.ResponseWriter, r *http.Request) error {
	var userID uuid.UUID
	// TODO: get user id from middleware
	//userID = middleware.GetUser(r)

	if userID == uuid.Nil {
		return api.Unauthorized()
	}

	ctx := r.Context()

	gotUser, err := h.userUsecase.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	api.WriteSuccess(ctx, w, gotUser)

	return nil
}
