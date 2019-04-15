package handler

import (
	"net/http"
	"regexp"
)

type Handler func(*Context)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type App struct {
	Routes        []Route
	DefaultRoutes map[string]Handler
}

func handle(pattern string, handler Handler) Route {
	re := regexp.MustCompile(pattern)
	return Route{Pattern: re, Handler: handler}
}

func (a *App) Handle(pattern string, handler Handler) {
	a.Routes = append(a.Routes, handle(pattern, handler))
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{Request: r, ResponseWriter: w}

	if customRoute(ctx, a) {
		return
	}
	a.DefaultRoutes["Route404"](ctx)
}

func AppHandler() *App {
	return &App{
		Routes: []Route{
			handle(`^/$`, RouteWiki),
			handle(`^/200$`, Route200),
			handle(`^/400$`, Route400),
			handle(`^/500$`, Route500),
		},
		DefaultRoutes: map[string]Handler{
			"Route200": Route200,
			"Route400": Route400,
			"Route404": Route404,
			"Route500": Route500,
		},
	}
}
