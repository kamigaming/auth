package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kamigaming/auth/internal/app/auth/ui"
)

// API version 1.
type APIv1 struct {
	login  ui.HTTPEntry
	signup ui.HTTPEntry
}

// Create new APIv1.
func NewAPIv1() APIv1 {
	return APIv1{
		login:  nil,
		signup: nil,
	}
}

func (api *APIv1) Route() http.Handler {
	router := chi.NewRouter()
	router.Mount("/login", api.login.Route())
	router.Mount("/signup", api.signup.Route())
	return router
}
