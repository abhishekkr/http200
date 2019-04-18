package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func customHandler(ctx *Context, handlerDetail HandlerDetail) Handler {
	return func(ctx *Context) {
		ctx.Text(handlerDetail.Status, handlerDetail.Body)
	}
}

func RoutePost(ctx *Context) {
	if ctx.Request.ContentLength == 0 {
		http.Error(ctx.ResponseWriter, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	ctx.Request.ParseForm()
	var route Route
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Text(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = json.Unmarshal(body, &route)
	if err != nil {
		ctx.Text(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ctx.App.Routes = append(ctx.App.Routes,
		handle(route.Method,
			route.PatternStr,
			customHandler(ctx, route.HandlerDetail)))
	ctx.Text(http.StatusOK, "New Route Added")
}

func customRoute(ctx *Context) bool {
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
