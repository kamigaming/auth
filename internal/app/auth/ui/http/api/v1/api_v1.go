package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// API version 1.
type APIv1 struct {
}

// Create new APIv1.
func NewAPIv1() APIv1 {
	return APIv1{}
}

func (api *APIv1) Route() http.Handler {
	router := chi.NewRouter()
	return router
}
