// Package httprouter provides convenient adapter the julienschmidt/httprouter
// for integrating with net/http and context
package httprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bbqtd/httprouter/params"
)

type Router struct {
	*httprouter.Router
}

func New() *Router {
	return &Router{
		httprouter.New(),
	}
}

func (r *Router) Get(path string, handler http.Handler)     { r.GET(path, wrapHandler(handler)) }
func (r *Router) Post(path string, handler http.Handler)    { r.POST(path, wrapHandler(handler)) }
func (r *Router) Delete(path string, handler http.Handler)  { r.DELETE(path, wrapHandler(handler)) }
func (r *Router) Patch(path string, handler http.Handler)   { r.PATCH(path, wrapHandler(handler)) }
func (r *Router) Put(path string, handler http.Handler)     { r.PUT(path, wrapHandler(handler)) }
func (r *Router) Options(path string, handler http.Handler) { r.OPTIONS(path, wrapHandler(handler)) }
func (r *Router) Head(path string, handler http.Handler)    { r.HEAD(path, wrapHandler(handler)) }

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r = params.NewRequestContext(r, ps)
		h.ServeHTTP(w, r)
	}
}
