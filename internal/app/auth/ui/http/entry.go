package http

import "net/http"

// HTTP Entry that can be routed to a handler.
type Entry interface {

	// Route this entry to an http.Handler.
	Route() http.Handler
}
