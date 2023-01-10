package vantuz

import (
	"fmt"
	"net/http"
	"strings"
)

type Logger struct {
	enabled bool
}

func (l *Logger) log(format string, a ...any) {
	if !l.enabled {
		return
	}
	var msg = fmt.Sprintf("[vantuz] "+format, a...)
	fmt.Println(msg)
}

func (l *Logger) request(req *http.Request) {
	if !l.enabled || req == nil {
		return
	}
	go func() {
		l.log("==== Request: %v ====", req.URL.String())
		for k, v := range req.Header {
			l.log("%s: %s", k, strings.Join(v, ","))
		}
	}()
}

func (l *Logger) response(resp *http.Response) {
	if !l.enabled || resp == nil {
		return
	}
	l.log("==== Response (%v): %v ====", resp.StatusCode, resp.Request.URL.String())
}

func (l *Logger) body(body []byte) {
	if !l.enabled {
		return
	}
	if body == nil {
		l.log("Body: nil")
		return
	}
	l.log("Body: %v", string(body))
}
