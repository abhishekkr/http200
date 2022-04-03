package httplog

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	emptyBody = bytes.NewBuffer([]byte(""))
)

func TestLogRequestMeta(t *testing.T) {
	req, err := http.NewRequest("XET", "/TestAppHandlerHandle", emptyBody)
	if err != nil {
		t.Fatal(err)
	}

	result := captureOutput(func() { logRequestMeta(req) })
	assert.True(t, strings.Contains(result, "---\nMethod: XET\n"))
}

func TestLogRequestHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/TestLog", emptyBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("x-this-use", "test")

	result := captureOutput(func() { logRequestHeaders(req) })
	assert.True(t, strings.Contains(result, "Headers:\n"))
	assert.True(t, strings.Contains(result, "X-This-Use: [test]"))
}

func TestLogRequestEncoding(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"status": "success"}`))
	req, err := http.NewRequest("POST", "/TestLog", body)
	if err != nil {
		t.Fatal(err)
	}
	req.TransferEncoding = []string{"chunked"}

	result := captureOutput(func() { logRequestEncoding(req) })
	assert.True(t, strings.Contains(result, "Encoding:\n"))
	assert.True(t, strings.Contains(result, "chunked"))
}

func TestLogRequestBody(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"status": "success"}`))
	req, err := http.NewRequest("POST", "/TestLog", body)
	if err != nil {
		t.Fatal(err)
	}

	result := captureOutput(func() { logRequestBody(req) })
	assert.True(t, strings.Contains(result, "Body:\n"))
	assert.True(t, strings.Contains(result, "success"))
}

func captureOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout, stderr := os.Stdout, os.Stderr
	defer func() {
		os.Stdout, os.Stderr = stdout, stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout, os.Stderr = w, w
	log.SetOutput(w)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, r)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	w.Close()
	return <-out
}
