package router

import (
	"github.com/brecode/urlLookup/handler"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Router struct {
	Deps
	*mux.Router
}

type Deps struct {
	Logger  *log.Logger
	Handler handler.API
}

// RouterOption defines the option function for Router
type RouterOption func(*Router)

// UseRouterDeps returns RouterOption that can inject custom dependencies.
func UseRouterDeps(rd func(*Deps)) RouterOption {
	return func(r *Router) {
		rd(&r.Deps)
	}
}

// NewRouter returns a router with routes defined in this package
func NewRouter(opts ...RouterOption) *Router {

	r := &Router{}
	for _, o := range opts {
		o(r)
	}

	return r
}

func (r *Router) Init() error {

	r.Router = mux.NewRouter().StrictSlash(true)

	// a list of routes to be registered
	for _, route := range r.GetRoutes() {
		r.Router.Methods(route.Method).
			PathPrefix(route.Pattern).
			Subrouter().
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return nil
}

func (r *Router) Close() error {
	return nil
}
