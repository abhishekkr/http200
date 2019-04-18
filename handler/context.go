package handler

import (
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
	App    *App
}

func (c *Context) Text(code int, body string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)

	io.WriteString(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}

func Route200(ctx *Context) {
	ctx.Text(http.StatusOK, "OK")
}

func Route400(ctx *Context) {
	ctx.Text(http.StatusBadRequest, "Bad Request")
}

func Route404(ctx *Context) {
	ctx.Text(http.StatusNotFound, "Not Found")
}

func Route500(ctx *Context) {
	ctx.Text(http.StatusInternalServerError, "Internal Server Error")
}
