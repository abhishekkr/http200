package httplog

import (
	"fmt"
	"net/http"
)

func LogRequest(req *http.Request) {
	logRequestMeta(req)
}

func logRequestMeta(req *http.Request) {
	fmt.Printf(`---
Method: %s
URL: %s
Protocol: %s
ContentLength: %d
Host: %s
RemoteAddr: %s
`, req.Method, req.URL, req.Proto,
		req.ContentLength, req.Host, req.RemoteAddr)
	logRequestEncoding(req)
	logRequestHeaders(req)
}

func logRequestHeaders(req *http.Request) {
	if len(req.Header) <= 0 {
		return
	}
	fmt.Println("Headers:")
	for header, val := range req.Header {
		fmt.Printf("  %s: %s\n", header, val)
	}
}

func logRequestEncoding(req *http.Request) {
	if len(req.TransferEncoding) <= 0 {
		return
	}
	fmt.Println("Encoding:")
	for encoding := range req.TransferEncoding {
		fmt.Printf("  %s\n", encoding)
	}
}
