package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Deps
}

type Deps struct {
	Logger *log.Logger
}

// HandlerOption defines the option function for Router
type HandlerOption func(*Handler)

// UseHandlerDeps returns HandlerOption that can inject custom dependencies.
func UseHandlerDeps(hd func(*Deps)) HandlerOption {
	return func(h *Handler) {
		hd(&h.Deps)
	}
}

// NewHandler returns a new Handler
func NewHandler(opts ...HandlerOption) *Handler {

	h := &Handler{}
	for _, o := range opts {
		o(h)
	}

	return h
}

// GetURLData will be the handling function for caller's GET requests
func (h *Handler) GetURLData() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// todo - implement the logic of a query here
		// todo - essentially query the db
	}
}

// UpdateURLData will be the handling function for caller's Update requests
func (h *Handler) UpdateURLData() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// todo - implement the logic of a POST request here
		// todo - save in the db
	}
}
