package handler

import (
	"net/http"
	"regexp"

	"github.com/abhishekkr/gol/golenv"
)

var (
	DefaultRoute = golenv.OverrideIfEnv("HTTP200_DEFAULT_ROUTE", "Route404")
)

type Handler func(*Context)

type HandlerDetail struct {
	Status int
	Body   string
}

type Route struct {
	Method        string
	PatternStr    string
	HandlerDetail HandlerDetail

	Pattern *regexp.Regexp
	Handler Handler
}

type App struct {
	Routes        []Route
	DefaultRoutes map[string]Handler
}

func handle(method string, pattern string, handler Handler) Route {
	re := regexp.MustCompile(pattern)
	return Route{
		Method:  method,
		Pattern: re,
		Handler: handler,
	}
}

func (a *App) Handle(method string, pattern string, handler Handler) {
	a.Routes = append(a.Routes, handle(method, pattern, handler))
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{Request: r, ResponseWriter: w, App: a}
	if customRoute(ctx) {
		return
	}
	a.DefaultRoutes[DefaultRoute](ctx)
}

func AppHandler() *App {
	return &App{
		Routes: []Route{
			handle("GET", `^/200$`, Route200),
			handle("GET", `^/400$`, Route400),
			handle("GET", `^/404$`, Route404),
			handle("GET", `^/500$`, Route500),
			handle("GET", `^/$`, RouteWiki),
			handle("POST", `^/$`, RoutePost),
		},
		DefaultRoutes: map[string]Handler{
			"Route200": Route200,
			"Route400": Route400,
			"Route404": Route404,
			"Route500": Route500,
		},
	}
}
