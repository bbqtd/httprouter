// Package params provides utility functions for using context with httprouter
// parameters
package params

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type key int

const paramsKey key = 0

func NewContext(ctx context.Context, params httprouter.Params) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}

func NewRequestContext(r *http.Request, params httprouter.Params) *http.Request {
	ctx := r.Context()
	ctx = NewContext(ctx, params)
	return r.WithContext(ctx)
}

func FromRequest(req *http.Request) (httprouter.Params, bool) {
	params, ok := req.Context().Value(paramsKey).(httprouter.Params)
	return params, ok
}

func FromContext(ctx context.Context) (httprouter.Params, bool) {
	params, ok := ctx.Value(paramsKey).(httprouter.Params)
	return params, ok
}
