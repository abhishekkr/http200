package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func customHandler(ctx *Context, handlerDetail HandlerDetail) Handler {
	return func(ctx *Context) {
		ctx.Text(handlerDetail.Status, handlerDetail.Body)
	}
}

func customRoute(ctx *Context) bool {
	log.Println(ctx.Method)
	for _, rt := range ctx.App.Routes {
		matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path)
		if len(matches) <= 0 || rt.Method != ctx.Method {
			continue
		}
		if len(matches) > 1 {
			ctx.Params = matches[1:]
		}
		rt.Handler(ctx)
		return true
	}
	return false
}

func RoutePost(ctx *Context) {
	if ctx.Request.ContentLength == 0 {
		http.Error(ctx.ResponseWriter, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(ctx.Request.Body)
	body := buf.String()
	log.Println("~~~", body)
	log.Println("~~~", len(body))

	var route Route
	err := json.NewDecoder(ctx.Request.Body).Decode(&route)
	if err != nil {
		ctx.Text(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.App.Routes = append(ctx.App.Routes,
		handle(route.Method,
			route.PatternStr,
			customHandler(ctx, route.HandlerDetail)))
	log.Println("~~~", route.Method)
	log.Println("~~~", route.PatternStr)
	log.Println("~~~", route.HandlerDetail)
	log.Println("~~~", ctx.Request.Header)
	ctx.Text(http.StatusOK, "New Route Added")
}
