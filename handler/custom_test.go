package handler

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutePost(t *testing.T) {
	app := &App{}
	app.Handle("POST", `^/TestMe$`, RoutePost)
	body := bytes.NewBuffer([]byte(`{"status": "success"}`))
	req, err := http.NewRequest("POST", "/TestMe", body)
	if err != nil {
		t.Fatal(err)
	}

	resp := NewRespWriter()
	ctx := &Context{resp, req, nil, app}
	RoutePost(ctx)
	s := string(resp.body)

	assert.Equal(t, resp.statusCode, 200)
	assert.True(t, strings.Contains(s, "New Route Added"))
}
