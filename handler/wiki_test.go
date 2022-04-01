package handler

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteWiki(t *testing.T) {
	resp := NewRespWriter()
	ctx := &Context{resp, nil, nil, nil}
	RouteWiki(ctx)
	s := string(resp.body)
	assert.True(t, strings.Contains(s, "<html"))
}
