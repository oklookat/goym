package vantuz

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestRateLimit(t *testing.T) {
	var expected = 5

	var reqCount = 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCount++
		if _, err := w.Write([]byte("hello")); err != nil {
			t.Fatal(err)
		}
	}))
	defer srv.Close()

	var cl = C().EnableDevMode()
	for i := 0; i < expected; i++ {
		_, err := cl.R().Get(context.Background(), srv.URL)
		if err != nil {
			t.Fatal(err)
		}
	}

	if reqCount != expected {
		t.Fatalf("expected %v requests, got: %v", expected, reqCount)
	}
}
