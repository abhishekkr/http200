package handler

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextText(t *testing.T) {
	resp := NewRespWriter()
	ctx := Context{resp, nil, nil, nil}
	ctx.Text(101, "Upgrade Protocol")
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 101)
	assert.True(t, strings.Contains(s, "Upgrade Protocol"))
}

func TestRoute200(t *testing.T) {
	resp := NewRespWriter()
	ctx := Context{resp, nil, nil, nil}
	Route200(&ctx)
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 200)
	assert.True(t, strings.Contains(s, "OK"))
}

func TestRoute400(t *testing.T) {
	resp := NewRespWriter()
	ctx := Context{resp, nil, nil, nil}
	Route400(&ctx)
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 400)
	assert.True(t, strings.Contains(s, "Bad Request"))
}

func TestRoute404(t *testing.T) {
	resp := NewRespWriter()
	ctx := Context{resp, nil, nil, nil}
	Route404(&ctx)
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 404)
	assert.True(t, strings.Contains(s, "Not Found"))
}

func TestRoute500(t *testing.T) {
	resp := NewRespWriter()
	ctx := Context{resp, nil, nil, nil}
	Route500(&ctx)
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 500)
	assert.True(t, strings.Contains(s, "Internal Server Error"))
}

type respWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewRespWriter() *respWriter {
	return &respWriter{
		header: http.Header{},
	}
}

func (w *respWriter) Header() http.Header {
	return w.header
}

func (w *respWriter) Write(b []byte) (int, error) {
	w.body = b
	return len(b), nil
}

func (w *respWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}
