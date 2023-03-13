package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/kamigaming/auth/internal/app/auth/ui/http/api/v1"
)

func main() {
	apiv1 := v1.NewAPIv1()
	router := chi.NewRouter()
	router.Mount("/api/v1", apiv1.Route())
	// @todo #2:30m Let users configure the address of the server.
	// @todo #2:30m Use https instead of http (make it configurable).
	http.ListenAndServe("localhost:8080", router)
}
