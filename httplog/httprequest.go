package httplog

import (
	"net/http"

	"github.com/abhishekkr/http200/persist"
)

func LogRequest(req *http.Request) {
	logRequestMeta(req)
	persist.InsertRequest(req)
}
