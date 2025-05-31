package fastgo

import (
	"net/http"
	"strings"
)

type Route struct {
	Method     string
	Path       string
	Handler    HandlerFunc
	PathParts  []string
	ParamNames []string
}

type Router struct {
	routes map[string][]*Route
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string][]*Route),
	}
}

func (r *Router) AddRoute(method string, path string, handler HandlerFunc) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var paramNames []string
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			paramNames = append(paramNames, p[1:])
		}
	}

	route := &Route{
		Method:     method,
		Path:       path,
		Handler:    handler,
		PathParts:  parts,
		ParamNames: paramNames,
	}
	r.routes[method] = append(r.routes[method], route)
}

func (r *Router) Get(path string, handler HandlerFunc) {
	r.AddRoute(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler HandlerFunc) {
	r.AddRoute(http.MethodPost, path, handler)
}

func (r *Router) ServeHTTP(ctx *Ctx) {
	reqParts := strings.Split(strings.Trim(ctx.Req.URL.Path, "/"), "/")
	routes := r.routes[ctx.Req.Method]

	for _, route := range routes {
		if len(reqParts) != len(route.PathParts) {
			continue
		}

		params := make(map[string]string)
		match := true
		for i, part := range route.PathParts {
			if strings.HasPrefix(part, ":") {
				params[part[1:]] = reqParts[i]
			} else if part != reqParts[i] {
				match = false
				break
			}
		}

		if match {
			ctx.Params = params
			route.Handler(ctx)
			return
		}
	}

	ctx.Next(HTTPErrorf(http.StatusNotFound, "Not Found"))
}
