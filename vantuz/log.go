package vantuz

import (
	"fmt"
	"io"
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
	l.log("==== %s: %v ====", req.Method, req.URL.String())
	for k, v := range req.Header {
		l.log("[header] %s: %s", k, strings.Join(v, ","))
	}
	if req.Body != nil {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, req.Body)
		if err == nil {
			l.log("[body] %s", buf.String())
		}
	}
}

func (l *Logger) response(resp *http.Response) {
	if !l.enabled || resp == nil {
		return
	}
	l.log("==== RESPONSE (%v): %v ====", resp.StatusCode, resp.Request.URL.String())
}

func (l *Logger) responseBody(body []byte) {
	if !l.enabled {
		return
	}
	var bodyStr = "nil"
	if len(body) != 0 {
		bodyStr = string(body)
	}
	l.log("[body]", bodyStr)
}
