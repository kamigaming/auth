package main

import (
	"net/http"

	v1 "github.com/kamigaming/auth/internal/app/auth/ui/http/api/v1"
)

func main() {
	apiv1 := v1.NewAPIv1()
	// @todo #2:30m Let users configure the address of the server.
	// @todo #2:30m Use https instead of http (make it configurable).
	http.ListenAndServe("localhost:8080", apiv1.Route())
}
