package api

import "net/http"

// Handler groups all HHTP handlers for the applications.
type Handler struct{}

// NewHandler returns a new HTTP handler instance.
func NewHandlers() *Handler {
	return &Handler{}
}

// RegisterRoutes attaches all API routes to the given ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// TODO: add your routes here in new projects
}
