package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request) error

func Handle(h HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		if err = h(w, r); err == nil {
			return
		}

		var ae Error
		if errors.As(err, &ae) {
			w.WriteHeader(ae.httpStatus)
			_ = json.NewEncoder(w).Encode(Response{
				Code:    ae.code,
				Message: ae.msg,
			})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(Response{
				Code:    ErrInternalServerError.code,
				Message: ErrInternalServerError.msg,
			})
		}
	}
}

func WriteSuccess(_ context.Context, w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Success(data))
}
