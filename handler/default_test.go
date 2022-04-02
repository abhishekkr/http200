package handler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	h := handle("X", `^Y`, nil)
	r := Route{Method: "X", Pattern: regexp.MustCompile(`^Y`), Handler: nil}
	assert.EqualValues(t, h, r)
}

func TestAppHandler(t *testing.T) {
	app := AppHandler()
	if assert.NotNil(t, app, "AppHandler gave nil") {
		assert.Greater(t, len(app.Routes), 0, "AppHandler returns empty app.Routes")
		assert.Greater(t, len(app.DefaultRoutes), 0, "AppHandler returns empty app.DefaultRoutes")
	}
}

func TestAppHandlerHandle(t *testing.T) {
	app := AppHandler()
	app.Handle("GET", `^/TestAppHandlerHandle$`, Route200)
	req, err := http.NewRequest("GET", "/TestAppHandlerHandle", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, checkRoute(app, req, http.StatusOK))
}

func TestAppHandlerServeHTTP(t *testing.T) {
	app := AppHandler()
	req, err := http.NewRequest("GET", "/200", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, checkRoute(app, req, http.StatusOK))
}

func checkRoute(app *App, req *http.Request, status int) bool {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.ServeHTTP)
	handler.ServeHTTP(rr, req)

	if rr.Code != status {
		return false
	}
	return true
}
